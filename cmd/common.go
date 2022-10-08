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
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/go-git/go-git/v5"
	"github.com/kyokomi/emoji/v2"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Revision     uint64              `yaml:"revision"`
	Ignores      []string            `yaml:"ignores"`
	Dependencies map[string][]string `yaml:"dependencies"`
}

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
	VersionUpdate      string
}

func getTags(r *git.Repository) ([]string, error) {
	tags := []string{}

	iter, err := r.Tags()
	if err != nil {
		return nil, err
	}

	for {
		ref, err := iter.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		tags = append(tags, ref.Name().Short())

	}
	return tags, nil
}

func ReplaceVersion(data string, appInfo *AppInfo, packageInfo *PackageInfo) string {
	version := packageInfo.Version
	data = strings.ReplaceAll(data, "{{{OS_FLAVOUR}}}", appInfo.OS_Flavour)
	data = strings.ReplaceAll(data, "{{{VERSION_ORIGINAL}}}", version.Original())
	data = strings.ReplaceAll(data, "{{{VERSION}}}", fmt.Sprintf("%v.%v.%v", version.Major(), version.Minor(), version.Patch()))
	data = strings.ReplaceAll(data, "{{{VERSION_MAJOR_MINOR}}}", fmt.Sprintf("%v.%v", version.Major(), version.Minor()))
	data = strings.ReplaceAll(data, "{{{VERSION_MAJOR}}}", fmt.Sprintf("%v", version.Major()))
	data = strings.ReplaceAll(data, "{{{VERSION_MINIOR}}}", fmt.Sprintf("%v", version.Minor()))
	data = strings.ReplaceAll(data, "{{{VERSION_PATCH}}}", fmt.Sprintf("%v", version.Patch()))
	data = strings.ReplaceAll(data, "{{{VERSION_PRERELEASE}}}", fmt.Sprintf("%v", version.Prerelease()))
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
			VersionUpdate:      FindPatchFile(packageInfo, "version/update.sh"),
		}
		patchInfo.PackageInfo = packageInfo

		patchInfos = append(patchInfos, patchInfo)
	}

	return patchInfos, nil
}

func packageSearch(n *parser.Node) {
	switch strings.ToLower(n.Value) {
	case "run":
		{
			command := n.Next.Value
			{
				packagesRegex, _ := regexp.Compile("component_unpack \"([^\"]*)\" \"([^\"]*)\"")
				packageSubmatchGroups := packagesRegex.FindAllStringSubmatch(command, -1)

				for _, groups := range packageSubmatchGroups {
					if len(groups) != 3 {
						log.Fatal("parse error")
					}
					packageName := groups[1]
					packageVersion, err := semver.NewVersion(groups[2])
					if err != nil {
						log.Fatal("parse error")
					}
					fmt.Println(packageName, packageVersion)
				}
			}

			{
				packagesRegex, _ := regexp.Compile("-f ([^ ]*)-([0-9]+.[0-9]+.[0-9]+-[^ ]+)-linux-[^ ]*[.]tar[.]gz \\]")
				packageSubmatchGroups := packagesRegex.FindAllStringSubmatch(command, -1)

				for _, groups := range packageSubmatchGroups {
					if len(groups) != 3 {
						log.Fatal("parse error")
					}
					packageName := groups[1]
					packageVersion, err := semver.NewVersion(groups[2])
					if err != nil {
						log.Fatal("parse error")
					}
					fmt.Println(packageName, packageVersion)
				}
			}
		}
	}
}

type PackageInstallCommand struct {
	Command        string
	PackageName    string
	PackageVersion *semver.Version
}

func packageInstallCommandSearch(n *parser.Node, a []PackageInstallCommand) []PackageInstallCommand {
	switch strings.ToLower(n.Value) {
	case "run":
		{
			command := n.Next.Value
			{
				packagesRegex, _ := regexp.Compile(`mkdir -p /tmp/bitnami/pkg/cache/ && cd /tmp/bitnami/pkg/cache/`)
				packageSubmatchGroups := packagesRegex.FindAllStringSubmatch(command, -1)
				if len(packageSubmatchGroups) > 0 {
					a = append(a, PackageInstallCommand{
						Command:     command,
						PackageName: "global",
					})
				}
			}
			{
				packagesRegex, _ := regexp.Compile("component_unpack \"([^\"]*)\" \"([^\"]*)\"")
				packageSubmatchGroups := packagesRegex.FindAllStringSubmatch(command, -1)

				for _, groups := range packageSubmatchGroups {
					if len(groups) != 3 {
						log.Fatal("parse error")
					}
					packageName := groups[1]
					packageVersion, err := semver.NewVersion(groups[2])
					if err != nil {
						log.Fatal("parse error")
					}
					a = append(a, PackageInstallCommand{
						Command:        command,
						PackageName:    packageName,
						PackageVersion: packageVersion,
					})
				}
			}

			{
				packagesRegex, _ := regexp.Compile("-f ([^ ]*)-([0-9]+.[0-9]+.[0-9]+-[^ ]+)-linux-[^ ]*[.]tar[.]gz \\]")
				packageSubmatchGroups := packagesRegex.FindAllStringSubmatch(command, -1)

				for _, groups := range packageSubmatchGroups {
					if len(groups) != 3 {
						log.Fatal("parse error")
					}
					packageName := groups[1]
					packageVersion, err := semver.NewVersion(groups[2])
					if err != nil {
						log.Fatal("parse error")
					}
					a = append(a, PackageInstallCommand{
						Command:        command,
						PackageName:    packageName,
						PackageVersion: packageVersion,
					})
				}
			}
		}
	}
	return a
}

func InspectDockerfile(path string) (*AppInfo, error) {
	var err error
	buf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Panic(err)
	}

	p := &Config{}
	err = yaml.Unmarshal(buf, p)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	dockerfile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	dockerfileString := string(dockerfile)

	dockerfileBuffer := bytes.NewBufferString(dockerfileString)

	result, err := parser.Parse(dockerfileBuffer)
	if err != nil {
		return nil, err
	}
	ast := result.AST
	nodes := []*parser.Node{ast}
	if ast.Children != nil {
		nodes = append(nodes, ast.Children...)
	}

	// version
	versionRegex, _ := regexp.Compile("(BITNAMI_IMAGE_VERSION|APP_VERSION)=\"([^\"]*)\"")
	versionSubmatchGroup := versionRegex.FindStringSubmatch(dockerfileString)

	packageNameRegex, _ := regexp.Compile("containers\\/bitnami\\/([^/\\\\]*)")
	appInfoName := packageNameRegex.FindStringSubmatch(strings.ReplaceAll(path, "\\", "/"))[1]

	var appVersion *semver.Version
	var OS_Arch string
	var OS_Flavour string
	var OS_Name string

	// from
	fromRegex, _ := regexp.Compile("FROM (.*)")
	fromSubmatchGroups := fromRegex.FindAllStringSubmatch(dockerfileString, -1)

	if len(fromSubmatchGroups) > 0 {
		groups := fromSubmatchGroups[len(fromSubmatchGroups)-1]
		switch groups[1] {
		case "scratch":
			OS_Arch = runtime.GOARCH
			OS_Flavour = "scratch"
			OS_Name = "linux"
			break
		case "bitnami/minideb:buster":
		case "docker.io/bitnami/minideb:buster":
			OS_Arch = runtime.GOARCH
			OS_Flavour = "debian-10"
			OS_Name = "linux"
			break
		case "bitnami/minideb:bullseye":
		case "docker.io/bitnami/minideb:bullseye":
			OS_Arch = runtime.GOARCH
			OS_Flavour = "debian-11"
			OS_Name = "linux"
			break
		}
	}

	// os
	osRegex, _ := regexp.Compile("(OS_ARCH|OS_FLAVOUR|OS_NAME)=\"([^\"]*)\"")
	osSubmatchsGroup := osRegex.FindAllStringSubmatch(dockerfileString, -1)

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
			Name: appInfoName,
		},
		OS_Arch,
		OS_Flavour,
		OS_Name,
		filepath.Dir(path),
		[]PackageInfo{},
	}

	// packages
	packages1 := []PackageInfo{}
	{
		packagesRegex, _ := regexp.Compile("component_unpack \"([^\"]*)\" \"([^\"]*)\"")
		packageSubmatchGroups := packagesRegex.FindAllStringSubmatch(dockerfileString, -1)

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
	}
	if len(packages1) == 0 {
		packagesRegex, _ := regexp.Compile(" -f ([^ ]*)-([0-9]+.[0-9]+.[0-9]+-[0-9]+)-linux-[^ ]*[.]tar[.]gz \\]")
		packageSubmatchGroups := packagesRegex.FindAllStringSubmatch(dockerfileString, -1)

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
		// packageVersionString := fmt.Sprintf("%v.%v.%v", packageInfo.Version.Major(), packageInfo.Version.Minor(), packageInfo.Version.Patch())
		// packageVersion, err := semver.NewVersion(packageVersionString)
		// if err != nil {
		// 	return nil, err
		// }
		// packageInfo.Version = packageVersion
		packages2 = append(packages2, packageInfo)
	}

	{
		packagesSorted := []PackageInfo{}
		for k, _ := range packages2 {
			pkgName := packages2[k].Name
			if val, ok := p.Dependencies[pkgName]; ok {
				idx := -1
				for k2, _ := range packagesSorted {
					idx = slices.Index(val, packagesSorted[k2].Name)
					if idx >= 0 {
						break
					}
				}
				fmt.Println(packages2, packagesSorted, packages2[k].Name, idx)
				if idx == 0 {
					packagesSorted = append([]PackageInfo{packages2[k]}, packagesSorted...)
				} else if idx >= 0 {
					tmp := []PackageInfo{}
					tmp = append(packagesSorted[0:idx], packages2[k])
					tmp = append(tmp, packagesSorted[:idx]...)
					packagesSorted = tmp
				} else {
					packagesSorted = append(packagesSorted, packages2[k])
				}
			} else {
				packagesSorted = append(packagesSorted, packages2[k])
			}
		}
		packages2 = packagesSorted
	}

	for _, p1 := range packages1 {
		for _, p2 := range packages2 {
			if p1.Name == p2.Name {
				p1.Version = p2.Version
			}
		}
	}

	if len(packages1) < len(packages2) {
		appInfo.Dependencies = packages2
	} else if len(packages1) > len(packages2) {
		appInfo.Dependencies = packages1
	} else if len(packages1) == len(packages2) {
		appInfo.Dependencies = packages1
	}

	if len(versionSubmatchGroup) != 3 {
		if len(appInfo.Dependencies) == 1 {
			appVersionString := appInfo.Dependencies[0].Version.String()
			appVersion, err = semver.NewVersion(appVersionString)
			if err != nil {
				return nil, err
			}
		} else {
			for _, p := range appInfo.Dependencies {
				fmt.Println(p.Name, p.Name == appInfoName)
			}
			fmt.Println(appInfo.OS_Flavour, appInfo.OS_Arch)
			return nil, errors.New("unknown version")
		}
	} else {
		appVersion, err = semver.NewVersion(strings.Split(versionSubmatchGroup[2], "-debian")[0])
		if err != nil {
			return nil, err
		}
	}
	if appVersion != nil {
		appVersionString := fmt.Sprintf("%v.%v.%v", appVersion.Major(), appVersion.Minor(), appVersion.Patch())
		appVersion, err = semver.NewVersion(appVersionString)
		if err != nil {
			return nil, err
		}
	}
	appInfo.Version = appVersion

	return appInfo, nil
}

func PatchDockerfile(appInfo *AppInfo) {
	var err error
	buf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Panic(err)
	}

	p := &Config{}
	err = yaml.Unmarshal(buf, p)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	scriptfiles, err := doublestar.FilepathGlob(filepath.Join(appInfo.Path, "**", "*.sh"))
	if err != nil {
		log.Panic(err)
	}

	for _, scriptfile := range scriptfiles {
		os.Chmod(scriptfile, 0755)
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
		if len(patchs) == 0 {
			patchFound = len(patchs) == 0
		}
	} else {
		patchFound = len(patchs) == 0
	}

	var golangBuilder bytes.Buffer
	var dockerfileFromBuilder bytes.Buffer
	var dockerfileInstallBuilder bytes.Buffer
	var newDockerfile bytes.Buffer

	if patchFound {
		for i, patch := range patchs {
			version := patch.PackageInfo.Version

			if patch.VersionUpdate != "" {
				var args []string

				err = os.Chmod(filepath.Join(wd, filepath.Dir(patch.VersionUpdate)), 0755)
				if err != nil {
					log.Panic(err)
				}

				err = os.Chmod(filepath.Join(wd, patch.VersionUpdate), 0755)
				if err != nil {
					log.Panic(err)
				}

				args = []string{
					"run", "--rm",
					"-e", fmt.Sprintf("PACKAGE=%v", patch.PackageInfo.Name),
					"-e", fmt.Sprintf("VERSION=%v.%v.%v", version.Major(), version.Minor(), version.Patch()),
					"-e", fmt.Sprintf("VERSION_MAJOR_MINOR=%v.%v", version.Major(), version.Minor()),
					"-e", fmt.Sprintf("VERSION_MAJOR=%v", version.Major()),
					"-e", fmt.Sprintf("VERSION_MINOR=%v", version.Minor()),
					"-e", fmt.Sprintf("VERSION_PATCH=%v", version.Patch()),
					"-e", fmt.Sprintf("VERSION_DIR=%v", "/work/prebuildfs/opt/bitnami/"),
					"-v", fmt.Sprintf("%v:/work", filepath.Join(wd, appInfo.Path)),
				}
				args = append(args, "-v", fmt.Sprintf("%v:/work/bash", filepath.Join(wd, filepath.Dir(patch.VersionUpdate))))

				args = append(args,
					"-w", "/work", "debian:bullseye-slim", "/bin/bash", "-c", strings.ReplaceAll(filepath.Join("bash", filepath.Base(patch.VersionUpdate)), "\\", "/"))

				cmd := exec.Command("docker", args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					log.Panic(err)
				}

				bitnamiComponentsBytes, err := ioutil.ReadFile(filepath.Join(appInfo.Path, "prebuildfs", "opt", "bitnami", ".bitnami_components.json"))
				if err != nil {
					log.Panic(err)
				}
				var bitnamiComponents map[string]interface{}

				err = json.Unmarshal(bitnamiComponentsBytes, &bitnamiComponents)
				if err != nil {
					log.Println(err)
					log.Panic(err)
				}

				NewVersion, err := semver.NewVersion(bitnamiComponents[patch.PackageInfo.Name].(map[string]interface{})["version"].(string))
				if err != nil {
					log.Panic(err)
				}
				patchs[i].PackageInfo.Version = NewVersion
			}
		}

		emoji.Println(fmt.Sprintf(":heavy_check_mark: %v:%v patch", appInfo.Name, appInfo.Version.Original()))

		originalDockerfile, err := ioutil.ReadFile(filepath.Join(appInfo.Path, "Dockerfile"))
		if err != nil {
			log.Panic(err)
		}

		dockerfileBuffer := bytes.NewBuffer(originalDockerfile)

		result, err := parser.Parse(dockerfileBuffer)
		if err != nil {
			log.Panic(err)
		}
		ast := result.AST
		nodes := []*parser.Node{ast}
		if ast.Children != nil {
			nodes = append(nodes, ast.Children...)
		}

		packageInstallCommands := []PackageInstallCommand{}
		for _, n := range nodes {
			packageInstallCommands = packageInstallCommandSearch(n, packageInstallCommands)
		}

		for _, patch := range patchs {
			emoji.Println("  ", patch)
			if patch.GolangBuild != "" {
				golangBuild, err := ioutil.ReadFile(patch.GolangBuild)
				if err != nil {
					log.Panic(err)
				}
				golangBuildString := string(golangBuild)
				golangBuildString = ReplaceVersion(golangBuildString, appInfo, &patch.PackageInfo)

				golangBuilder.Write([]byte(golangBuildString))
				golangBuilder.Write([]byte("\n\n"))
			}

			if patch.DockerFromPatch != "" {
				dockerFromPatch, err := ioutil.ReadFile(patch.DockerFromPatch)
				if err != nil {
					log.Panic(err)
				}
				dockerFromPatchString := string(dockerFromPatch)
				dockerFromPatchString = ReplaceVersion(dockerFromPatchString, appInfo, &patch.PackageInfo)

				dockerfileFromBuilder.Write([]byte(dockerFromPatchString))
				dockerfileFromBuilder.Write([]byte("\n\n"))
			}

			if patch.DockerInstallPatch != "" {
				dockerInstallPatch, err := ioutil.ReadFile(patch.DockerInstallPatch)
				if err != nil {
					log.Panic(err)
				}
				dockerInstallPatchString := string(dockerInstallPatch)
				dockerInstallPatchString = ReplaceVersion(dockerInstallPatchString, appInfo, &patch.PackageInfo)

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

			// sampleRegexp := regexp.MustCompile(`RUN mkdir -p /tmp/bitnami/pkg/cache/ && cd /tmp/bitnami/pkg/cache/ && .*tar.gz$`)
			// result := sampleRegexp.ReplaceAllString(originalDockerfileString, "# $0")
			// fmt.Println(string(result))

			for _, installCmd := range packageInstallCommands {
				spaces := regexp.MustCompile(`( )(   [ ]*)`)
				// fmt.Println(strings.ReplaceAll(installCmd.Command, "    ", " \\\n   "))
				if installCmd.PackageVersion != nil {
					originalDockerfileString = strings.ReplaceAll(originalDockerfileString,
						spaces.ReplaceAllString(installCmd.Command, " \\\n$2"),
						fmt.Sprintf("echo install %v %v", installCmd.PackageName, installCmd.PackageVersion.String()))
				} else {
					originalDockerfileString = strings.ReplaceAll(originalDockerfileString,
						spaces.ReplaceAllString(installCmd.Command, " \\\n$2"),
						fmt.Sprintf("echo install %v", installCmd.PackageName))
				}
			}

			versionSemver := fmt.Sprintf("%v.%v.%v", appInfo.Version.Major(), appInfo.Version.Minor(), appInfo.Version.Patch())
			replaceVars := map[string]string{
				"org.opencontainers.image.authors":     "https://github.com/ZCube/bitnami-compat",
				"org.opencontainers.image.description": "Application repackaged by ZCube for arm64 compatible, originaly Bitnami",
				"org.opencontainers.image.ref.name":    fmt.Sprintf("%v-%v-r%v", versionSemver, appInfo.OS_Flavour, p.Revision),
				"org.opencontainers.image.source":      "https://github.com/ZCube/bitnami-compat",
				"org.opencontainers.image.vendor":      "ZCube",
				"APP_VERSION":                          versionSemver,
			}

			{
				varsRegexp := regexp.MustCompile(`([^ \t]+)[ \t]*=[ \t]*"([^"]*)"`)
				vars := varsRegexp.FindAllStringSubmatch(originalDockerfileString, -1)
				for i, _ := range vars {
					if _, ok := replaceVars[vars[i][1]]; ok {
						originalDockerfileString = strings.ReplaceAll(originalDockerfileString,
							vars[i][0],
							fmt.Sprintf("%v=\"%v\"", vars[i][1], replaceVars[vars[i][1]]))
					}
				}
			}

			seperators := []string{
				"RUN apt-get autoremove --purge -y ",
				"RUN apt-get update && apt-get upgrade -y && \\",
				"FROM scratch",
			}
			var seperator string
			for _, v := range seperators {
				if strings.Contains(originalDockerfileString, v) {
					seperator = v
				}
			}
			if golangBuilder.Len() > 0 {
				originalDockerfileString = strings.ReplaceAll(originalDockerfileString,
					seperator,
					"COPY --from=golang-builder /opt/bitnami/ /opt/bitnami/ \n"+
						string(dockerfileInstallBuilder.Bytes())+
						"\n\n"+seperator)

			} else {
				originalDockerfileString = strings.Replace(originalDockerfileString,
					seperator,
					string(dockerfileInstallBuilder.Bytes())+
						"\n\n"+seperator, 1)
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

			if patch.BashPatch != "" {
				var args []string

				err = os.Chmod(filepath.Join(wd, filepath.Dir(patch.BashPatch)), 0755)
				if err != nil {
					log.Panic(err)
				}

				err = os.Chmod(filepath.Join(wd, patch.BashPatch), 0755)
				if err != nil {
					log.Panic(err)
				}

				args = []string{
					"run", "--rm",
					"-e", fmt.Sprintf("PACKAGE=%v", patch.PackageInfo.Name),
					"-e", fmt.Sprintf("VERSION=%v.%v.%v", version.Major(), version.Minor(), version.Patch()),
					"-e", fmt.Sprintf("VERSION_MAJOR_MINOR=%v.%v", version.Major(), version.Minor()),
					"-e", fmt.Sprintf("VERSION_MAJOR=%v", version.Major()),
					"-e", fmt.Sprintf("VERSION_MINOR=%v", version.Minor()),
					"-e", fmt.Sprintf("VERSION_PATCH=%v", version.Patch()),
					"-v", fmt.Sprintf("%v:/work", filepath.Join(wd, appInfo.Path)),
				}
				args = append(args, "-v", fmt.Sprintf("%v:/work/bash", filepath.Join(wd, filepath.Dir(patch.BashPatch))))

				args = append(args,
					"-w", "/work", "debian:bullseye-slim", strings.ReplaceAll(filepath.Join("bash", filepath.Base(patch.BashPatch)), "\\", "/"))

				cmd := exec.Command("docker", args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					log.Panic(err)
				}
			}
		}
	}
}
