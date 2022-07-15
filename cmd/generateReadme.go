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
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/cbroglie/mustache"
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// generateReadmeCmd represents the generateReadme command
var generateReadmeCmd = &cobra.Command{
	Use:   "generateReadme",
	Short: "generateReadme",
	Long:  `generateReadme`,
	Run: func(cmd *cobra.Command, args []string) {
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

		READMEHeader, err := ioutil.ReadFile("scripts/README-header.md")
		if err != nil {
			log.Panic(err)
		}

		READMEBody, err := ioutil.ReadFile("scripts/README-body.md")
		if err != nil {
			log.Panic(err)
		}

		READMETail, err := ioutil.ReadFile("scripts/README-tail.md")
		if err != nil {
			log.Panic(err)
		}

		var README bytes.Buffer

		_, err = README.Write(READMEHeader)
		if err != nil {
			log.Panic(err)
		}

		var dockerfiles []string
		if len(app) > 0 {
			dockerfiles, err = doublestar.FilepathGlob(fmt.Sprintf("bitnami-dockers/bitnami-docker-%v/**/Dockerfile", app))
		} else {
			dockerfiles, err = doublestar.FilepathGlob(fmt.Sprintf("bitnami-dockers/bitnami-docker-*/**/Dockerfile"))
		}
		if err != nil {
			log.Panic(err)
		}

		for i := range dockerfiles {
			if appInfo, err := InspectDockerfile(dockerfiles[i]); err == nil {
				// emoji.Println(appInfo.Dependencies)

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
					emoji.Println(fmt.Sprintf(":heavy_check_mark: %v:%v", appInfo.Name, appInfo.Version.Original()))
					READMEBodyString := string(READMEBody)

					var version string
					for _, path := range strings.Split(strings.ReplaceAll(appInfo.Path, "\\", "/"), "/") {
						if path == appInfo.OS_Flavour {
							break
						}
						version = path
					}
					READMEBodyString, err = mustache.Render(READMEBodyString,
						map[string]string{
							"APP":           appInfo.Name,
							"VERSION":       appInfo.Version.Original(),
							"VERSION_MAJOR": version,
							"OS_FLAVOUR":    appInfo.OS_Flavour,
							"OS_NAME":       appInfo.OS_Name,
							"REVISION":      fmt.Sprint(p.Revision),
						})
					if err != nil {
						log.Panic(err)
					}
					_, err = README.Write([]byte(READMEBodyString))
					if err != nil {
						log.Panic(err)
					}
				}
			}
		}

		_, err = README.Write(READMETail)
		if err != nil {
			log.Panic(err)
		}

		err = ioutil.WriteFile("README.md", README.Bytes(), 0644)
		if err != nil {
			log.Panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateReadmeCmd)
	generateReadmeCmd.PersistentFlags().StringVar(&app, "app", "", "app")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateReadmeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateReadmeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
