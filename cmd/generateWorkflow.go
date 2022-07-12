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
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// generateWorkflowCmd represents the generateWorkflow command
var generateWorkflowCmd = &cobra.Command{
	Use:   "generateWorkflow",
	Short: "generateWorkflow",
	Long:  `generateWorkflow`,
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

		buildOnPushHeader, err := ioutil.ReadFile("scripts/build-on-push-header")
		if err != nil {
			log.Panic(err)
		}

		buildOnPushBody, err := ioutil.ReadFile("scripts/build-on-push-body")
		if err != nil {
			log.Panic(err)
		}

		buildOnPushTail, err := ioutil.ReadFile("scripts/build-on-push-tail")
		if err != nil {
			log.Panic(err)
		}

		var buildOnPush bytes.Buffer

		_, err = buildOnPush.Write(buildOnPushHeader)
		if err != nil {
			log.Panic(err)
		}

		dockerfiles, err := doublestar.FilepathGlob("bitnami-docker-*/**/Dockerfile")

		if err != nil {
			log.Panic(err)
		}

		for i := range dockerfiles {
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
					patchFound = false
				}

				if patchFound {
					fmt.Println(fmt.Sprintf("(o) %v:%v", appInfo.Name, appInfo.Version.Original()))
					buildOnPushBodyString := string(buildOnPushBody)
					version := strings.Split(strings.ReplaceAll(appInfo.Path, "\\", "/"), "/")[1]
					buildOnPushBodyString, err = mustache.Render(buildOnPushBodyString,
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
					_, err = buildOnPush.Write([]byte(buildOnPushBodyString))
					if err != nil {
						log.Panic(err)
					}
				}
			}
		}

		_, err = buildOnPush.Write(buildOnPushTail)
		if err != nil {
			log.Panic(err)
		}

		err = ioutil.WriteFile(".github/workflows/build-on-push.yml", buildOnPush.Bytes(), 0644)
		if err != nil {
			log.Panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateWorkflowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateWorkflowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateWorkflowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
