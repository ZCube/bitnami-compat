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
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

var tagTo string
var destCreds string

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push",
	Long:  `push`,
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

		var dockerfiles []string
		if len(app) > 0 {
			dockerfiles, err = doublestar.FilepathGlob(fmt.Sprintf("containers/bitnami/%v/**/Dockerfile", app))
		} else {
			dockerfiles, err = doublestar.FilepathGlob(fmt.Sprintf("containers/bitnami/*/**/Dockerfile"))
		}
		if err != nil {
			log.Panic(err)
		}

		var lastAppInfo *AppInfo
		lastAppInfo = nil

		for i := range dockerfiles {
			emoji.Println(dockerfiles[i])
			if appInfo, err := InspectDockerfile(dockerfiles[i]); err == nil {
				if slices.Contains(p.Ignores, appInfo.Name) {
					continue
				}

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
					if len(patchs) == 0 {
						patchFound = len(patchs) == 0
					}
				} else {
					patchFound = len(patchs) == 0
				}

				if patchFound {
					var args []string
					{
						var version string
						for _, path := range strings.Split(strings.ReplaceAll(appInfo.Path, "\\", "/"), "/") {
							if path == appInfo.OS_Flavour {
								break
							}
							version = path
						}
						versionSemver := fmt.Sprintf("%v.%v.%v", appInfo.Version.Major(), appInfo.Version.Minor(), appInfo.Version.Patch())

						// skopeo copy --all --dest-creds=zcube:${{ secrets.DOCKER_HUB_KEY }}
						// docker://${{ env.REGISTRY }}/${{ steps.repository.outputs.lowercase }}/${{ matrix.APP }}
						// docker://zcube/bitnami-compat-${{ matrix.APP }}:${{ matrix.VERSION }}
						args = []string{
							"copy", "--all",
							"--dest-creds", destCreds,
							fmt.Sprintf("docker://%v/%v:%v-%v-r%v", tag, appInfo.Name, versionSemver, appInfo.OS_Flavour, p.Revision),
						}

						dests := []string{
							fmt.Sprintf("docker://%v%v:%v-%v-r%v", tagTo, appInfo.Name, version, appInfo.OS_Flavour, p.Revision),
							fmt.Sprintf("docker://%v%v:%v-%v", tagTo, appInfo.Name, version, appInfo.OS_Flavour),
							fmt.Sprintf("docker://%v%v:%v", tagTo, appInfo.Name, version),
							fmt.Sprintf("docker://%v%v:%v-%v-r%v", tagTo, appInfo.Name, versionSemver, appInfo.OS_Flavour, p.Revision),
							fmt.Sprintf("docker://%v%v:%v-%v", tagTo, appInfo.Name, versionSemver, appInfo.OS_Flavour),
							fmt.Sprintf("docker://%v%v:%v", tagTo, appInfo.Name, versionSemver),
						}

						for _, dest := range dests {
							args_ := slices.Clone(args)
							args_ = append(args_, dest)

							emoji.Println(args_)

							cmd := exec.Command("skopeo", args_...)
							cmd.Stdout = os.Stdout
							cmd.Stderr = os.Stderr
							if err := cmd.Run(); err != nil {
								log.Panic(err)
							}
						}

						lastAppInfo = appInfo
					}
				}
			} else {
				log.Panic(err)
			}
		}
		if lastAppInfo != nil {
			appInfo := lastAppInfo

			var args []string
			{
				versionSemver := fmt.Sprintf("%v.%v.%v", appInfo.Version.Major(), appInfo.Version.Minor(), appInfo.Version.Patch())

				args = []string{
					"copy", "--all",
					"--dest-creds", destCreds,
					fmt.Sprintf("docker://%v/%v:%v-%v-r%v", tag, appInfo.Name, versionSemver, appInfo.OS_Flavour, p.Revision),
				}

				dests := []string{
					fmt.Sprintf("docker://%v%v:latest", tagTo, appInfo.Name),
				}

				for _, dest := range dests {
					args_ := slices.Clone(args)
					args_ = append(args_, dest)

					emoji.Println(args_)

					cmd := exec.Command("skopeo", args_...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						log.Panic(err)
					}
				}

				lastAppInfo = appInfo
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.PersistentFlags().StringVar(&app, "app", "", "app")
	pushCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "ghcr.io/zcube/bitnami-compat", "tag")
	pushCmd.PersistentFlags().StringVar(&tagTo, "tag-to", "docker.io/zcube/bitnami-compat-", "tag-to")
	pushCmd.PersistentFlags().StringVar(&destCreds, "dest-creds", "", "dest-creds")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
