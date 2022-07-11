/*
Copyright © 2022 ZCube <zcubekr@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/bmatcuk/doublestar/v4"
)

type PackageInfo struct {
	Name    string
	Version *semver.Version
}

type AppInfo struct {
	PackageInfo
	OS_Arch      string
	OS_Flavour   string
	OS_Name      string
	Path         string
	Dependencies []PackageInfo
}

type PatchInfo struct {
	PackageInfo        PackageInfo
	BashPatch          string
	GolangBuild        string
	DockerFromPatch    string
	DockerInstallPatch string
}

func ReplaceVersion(data string, packageInfo PackageInfo) string {
	version := packageInfo.Version
	data = strings.ReplaceAll(data, "{{{VERSION}}}", fmt.Sprintf("%v.%v.%v", version.Major(), version.Minor(), version.Patch()))
	data = strings.ReplaceAll(data, "{{{VERSION_MAJOR_MINOR}}}", fmt.Sprintf("%v.%v", version.Major(), version.Minor()))
	data = strings.ReplaceAll(data, "{{{VERSION_MAJOR}}}", fmt.Sprintf("%v", version.Major()))
	return data
}

func FindPatchFile(packageInfo PackageInfo, filePath string) string {
	var paths []string
	version := packageInfo.Version

	// bash
	paths = []string{
		// fmt.Sprintf("patches/%v/%v.%v.%v-%v/%v", packageInfo.Name, version.Major(), version.Minor(), version.Patch(), version.Prerelease(), filePath),
		// fmt.Sprintf("patches/%v/%v.%v.%v/%v", packageInfo.Name, version.Major(), version.Minor(), version.Patch(), filePath),
		fmt.Sprintf("patches/%v/%v.%v/%v", packageInfo.Name, version.Major(), version.Minor(), filePath),
		fmt.Sprintf("patches/%v/%v/%v", packageInfo.Name, version.Major(), filePath),
	}

	for _, path := range paths {
		list, err := doublestar.FilepathGlob(path)
		if err != nil {
			continue
		}
		if len(list) > 0 {
			return list[0]
		}
	}
	return ""
}

func FindPatchs(appInfo *AppInfo) ([]PatchInfo, error) {
	var patchInfos []PatchInfo
	for _, packageInfo := range appInfo.Dependencies {
		var patchInfo PatchInfo

		patchInfo = PatchInfo{
			BashPatch:          FindPatchFile(packageInfo, "bash/install.sh"),
			GolangBuild:        FindPatchFile(packageInfo, "golang/install.sh"),
			DockerFromPatch:    FindPatchFile(packageInfo, "docker/Dockerfile.from"),
			DockerInstallPatch: FindPatchFile(packageInfo, "docker/Dockerfile.install"),
		}
		patchInfo.PackageInfo = packageInfo

		patchInfos = append(patchInfos, patchInfo)
	}

	return patchInfos, nil
}

func InspectDockerfile(path string) (*AppInfo, error) {
	dockerfile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	dockerfileString := string(dockerfile)

	// version
	versionRegex, _ := regexp.Compile("(BITNAMI_IMAGE_VERSION|APP_VERSION)=\"([^\"]*)\"")
	versionSubmatchGroup := versionRegex.FindStringSubmatch(dockerfileString)

	if len(versionSubmatchGroup) != 3 {
		return nil, errors.New("version not found")
	}

	appVersion, err := semver.NewVersion(strings.Split(versionSubmatchGroup[2], "-debian")[0])
	if err != nil {
		return nil, err
	}

	// os
	osRegex, _ := regexp.Compile("(OS_ARCH|OS_FLAVOUR|OS_NAME)=\"([^\"]*)\"")
	osSubmatchsGroup := osRegex.FindAllStringSubmatch(dockerfileString, -1)

	var OS_Arch string
	var OS_Flavour string
	var OS_Name string

	for _, groups := range osSubmatchsGroup {
		if len(groups) != 3 {
			continue
		}
		switch groups[1] {
		case "OS_ARCH":
			OS_Arch = groups[2]
			break
		case "OS_FLAVOUR":
			OS_Flavour = groups[2]
			break
		case "OS_NAME":
			OS_Name = groups[2]
			break
		}
	}

	appInfo := &AppInfo{
		PackageInfo{
			Version: appVersion,
		},
		OS_Arch,
		OS_Flavour,
		OS_Name,
		filepath.Dir(path),
		[]PackageInfo{},
	}

	appInfo.Name = strings.Split(strings.Split(strings.Join(strings.Split(path, "-")[2:], "-"), "/")[0], "\\")[0]

	// packages
	packagesRegex, _ := regexp.Compile("component_unpack \"([^\"]*)\" \"([^\"]*)\"")
	packageSubmatchGroups := packagesRegex.FindAllStringSubmatch(dockerfileString, -1)

	packages1 := []PackageInfo{}
	for _, groups := range packageSubmatchGroups {
		if len(groups) != 3 {
			return nil, errors.New("version not found")
		}
		packageName := groups[1]
		packageVersion, err := semver.NewVersion(groups[2])
		if err != nil {
			return nil, err
		}
		packages1 = append(packages1, PackageInfo{
			Name:    packageName,
			Version: packageVersion,
		})
	}

	bitnamiComponentsBytes, err := ioutil.ReadFile(filepath.Join(filepath.Dir(path), "prebuildfs", "opt", "bitnami", ".bitnami_components.json"))
	if err != nil {
		return nil, err
	}
	var bitnamiComponents map[string]interface{}

	err = json.Unmarshal(bitnamiComponentsBytes, &bitnamiComponents)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	packages2 := []PackageInfo{}
	for i := range bitnamiComponents {
		var packageInfo PackageInfo
		packageInfo.Name = i
		packageInfo.Version, err = semver.NewVersion(bitnamiComponents[i].(map[string]interface{})["version"].(string))
		if err != nil {
			return nil, err
		}
		packages2 = append(packages2, packageInfo)
	}

	if len(packages1) < len(packages2) {
		appInfo.Dependencies = packages2
	} else if len(packages1) > len(packages2) {
		appInfo.Dependencies = packages1
	} else if len(packages1) == len(packages2) {
		appInfo.Dependencies = packages1
	}
	return appInfo, nil
}

func PatchDockerfile(appInfo *AppInfo) {
	var err error
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	golangDockerfileHead, err := ioutil.ReadFile("scripts/Dockerfile.head")
	if err != nil {
		log.Panic(err)
	}

	patchFound := false
	var patchs []PatchInfo
	if patchs, err = FindPatchs(appInfo); err == nil {
		for _, patch := range patchs {
			patchFound = (patch.BashPatch != "" ||
				patch.DockerFromPatch != "" ||
				patch.DockerInstallPatch != "" ||
				patch.GolangBuild != "")
			if patchFound == false {
				break
			}
		}
	} else {
		patchFound = false
	}

	var golangBuilder bytes.Buffer
	var dockerfileFromBuilder bytes.Buffer
	var dockerfileInstallBuilder bytes.Buffer
	var newDockerfile bytes.Buffer

	if patchFound {
		fmt.Println(fmt.Sprintf("(o) %v:%v patch", appInfo.Name, appInfo.Version.Original()))

		originalDockerfile, err := ioutil.ReadFile(filepath.Join(appInfo.Path, "Dockerfile"))
		if err != nil {
			log.Panic(err)
		}
		for _, patch := range patchs {
			fmt.Println(patch)
			if patch.GolangBuild != "" {
				golangBuild, err := ioutil.ReadFile(patch.GolangBuild)
				if err != nil {
					log.Panic(err)
				}
				golangBuildString := string(golangBuild)
				golangBuildString = ReplaceVersion(golangBuildString, patch.PackageInfo)

				golangBuilder.Write([]byte(golangBuildString))
				golangBuilder.Write([]byte("\n\n"))
			}

			if patch.DockerFromPatch != "" {
				dockerFromPatch, err := ioutil.ReadFile(patch.DockerFromPatch)
				if err != nil {
					log.Panic(err)
				}
				dockerFromPatchString := string(dockerFromPatch)
				dockerFromPatchString = ReplaceVersion(dockerFromPatchString, patch.PackageInfo)

				dockerfileFromBuilder.Write([]byte(dockerFromPatchString))
				dockerfileFromBuilder.Write([]byte("\n\n"))
			}

			if patch.DockerInstallPatch != "" {
				dockerInstallPatch, err := ioutil.ReadFile(patch.DockerInstallPatch)
				if err != nil {
					log.Panic(err)
				}
				dockerInstallPatchString := string(dockerInstallPatch)
				dockerInstallPatchString = ReplaceVersion(dockerInstallPatchString, patch.PackageInfo)

				dockerfileInstallBuilder.Write([]byte(dockerInstallPatchString))
				dockerfileInstallBuilder.Write([]byte("\n\n"))
			}
		}

		if golangBuilder.Len() > 0 {
			newDockerfile.Write(golangDockerfileHead)
			newDockerfile.Write([]byte("\n"))
			ioutil.WriteFile(filepath.Join(appInfo.Path, "packages.sh"), golangBuilder.Bytes(), 0644)
		}

		{
			newDockerfile.Write(dockerfileFromBuilder.Bytes())
			// sed -i 's/RUN \. \/opt\/bitnami\/scripts\/libcomponent.sh/#RUN . \/opt\/bitnami\/scripts\/libcomponent.sh/g' _Dockerfile
			originalDockerfileString := string(originalDockerfile)

			originalDockerfileString = strings.ReplaceAll(originalDockerfileString,
				"RUN . /opt/bitnami/scripts/libcomponent.sh ",
				"# RUN . /opt/bitnami/scripts/libcomponent.sh ")

			if golangBuilder.Len() > 0 {
				originalDockerfileString = strings.ReplaceAll(originalDockerfileString,
					"RUN apt-get update && apt-get upgrade -y && \\",
					"COPY --from=golang-builder /opt/bitnami/ /opt/bitnami/ \n"+
						string(dockerfileInstallBuilder.Bytes())+
						"\n\nRUN apt-get update && apt-get upgrade -y && \\")

			} else {
				originalDockerfileString = strings.ReplaceAll(originalDockerfileString,
					"RUN apt-get update && apt-get upgrade -y && \\",
					string(dockerfileInstallBuilder.Bytes())+
						"\n\nRUN apt-get update && apt-get upgrade -y && \\")
			}

			if strings.Contains(originalDockerfileString, "bullseye") &&
				strings.Contains(originalDockerfileString, "buster") {
				log.Panic(errors.New("buster + bullseye"))
			}
			if strings.Contains(originalDockerfileString, "bullseye") &&
				strings.Contains(originalDockerfileString, "pgdg100") {
				log.Panic(errors.New("buster + bullseye"))
			}
			newDockerfile.Write([]byte(originalDockerfileString))
		}
		ioutil.WriteFile(filepath.Join(appInfo.Path, "Dockerfile.arm64"), newDockerfile.Bytes(), 0644)

		for _, patch := range patchs {
			version := patch.PackageInfo.Version

			err = os.Chmod(filepath.Join(wd, filepath.Dir(patch.BashPatch)), 0755)
			if err != nil {
				log.Panic(err)
			}

			var args []string
			if patch.BashPatch != "" {
				args = []string{
					"run", "--rm",
					"-e", fmt.Sprintf("VERSION=%v.%v.%v", version.Major(), version.Minor(), version.Patch()),
					"-e", fmt.Sprintf("VERSION_MAJOR_MINOR:%v.%v", version.Major(), version.Minor()),
					"-e", fmt.Sprintf("VERSION_MAJOR=%v", version.Major()),
					"-v", fmt.Sprintf("%v:/work", filepath.Join(wd, appInfo.Path)),
				}
				args = append(args, "-v", fmt.Sprintf("%v:/work/bash", filepath.Join(wd, filepath.Dir(patch.BashPatch))))

				args = append(args,
					"-w", "/work", "debian:bullseye-slim", strings.ReplaceAll(filepath.Join("bash", filepath.Base(patch.BashPatch)), "\\", "/"))

				cmd := exec.Command("docker", args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
