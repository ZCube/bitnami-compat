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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var tag string
var cacheFrom string
var cacheTo string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build",
	Long:  `build`,
	Run: func(cmd *cobra.Command, args []string) {
		buf, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Panic(err)
		}

		p := &Config{}
		err = yaml.Unmarshal(buf, p)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}

		fmt.Println(cmd.Flag("docker").Value.String())
		docker := cmd.Flag("docker").Value.String()
		fmt.Println(cacheFrom)
		fmt.Println(cacheTo)
		fmt.Println(tag)

		dockerfiles, err := doublestar.FilepathGlob(fmt.Sprintf("bitnami-dockers/bitnami-docker-%v/**/Dockerfile", docker))
		if err != nil {
			log.Panic(err)
		}

		wd, err := os.Getwd()
		if err != nil {
			log.Panic(err)
		}

		for i := range dockerfiles {
			fmt.Println(dockerfiles[i])
			if appInfo, err := InspectDockerfile(dockerfiles[i]); err == nil {
				// fmt.Println(appInfo.Dependencies)

				patchFound := false
				var err error
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
					patchFound = len(patchs) == 0
				}

				if patchFound {
					var args []string
					{
						args = []string{
							"buildx", "build",
						}

						version := strings.Split(strings.ReplaceAll(appInfo.Path, "\\", "/"), "/")[1]
						versionSemver := fmt.Sprintf("%v.%v.%v", appInfo.Version.Major(), appInfo.Version.Minor(), appInfo.Version.Patch())
						args = append(args, "-t", fmt.Sprintf("%v/%v:%v-%v-r%v", tag, appInfo.Name, version, appInfo.OS_Flavour, p.Revision))
						args = append(args, "-t", fmt.Sprintf("%v/%v:%v-%v", tag, appInfo.Name, version, appInfo.OS_Flavour))
						args = append(args, "-t", fmt.Sprintf("%v/%v:%v", tag, appInfo.Name, version))
						args = append(args, "-t", fmt.Sprintf("%v/%v:%v-%v-r%v", tag, appInfo.Name, versionSemver, appInfo.OS_Flavour, p.Revision))
						args = append(args, "-t", fmt.Sprintf("%v/%v:%v-%v", tag, appInfo.Name, versionSemver, appInfo.OS_Flavour))
						args = append(args, "-t", fmt.Sprintf("%v/%v:%v", tag, appInfo.Name, versionSemver))

						args = append(args,
							"-f", fmt.Sprintf(filepath.Join(wd, filepath.Dir(dockerfiles[i]), "Dockerfile.arm64")),
							fmt.Sprintf(filepath.Join(wd, filepath.Dir(dockerfiles[i]))),
						)

						fmt.Println(args)

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
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.PersistentFlags().String("docker", "", "docker")
	buildCmd.PersistentFlags().StringVar(&cacheFrom, "cache-from", "", "cache-from")
	buildCmd.PersistentFlags().StringVar(&cacheTo, "cache-to", "", "cache-to")
	buildCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "ghcr.io/zcube/bitnami-compat", "tag")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
