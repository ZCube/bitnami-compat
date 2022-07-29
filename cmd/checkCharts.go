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
	"regexp"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"
)

var level int

// checkChartsCmd represents the checkCharts command
var checkChartsCmd = &cobra.Command{
	Use:   "checkCharts",
	Short: "checkCharts",
	Long:  `checkCharts`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var dockerfiles []string
		dockerfiles, err = doublestar.FilepathGlob(fmt.Sprintf("containers/bitnami/*/**/Dockerfile"))
		if err != nil {
			log.Panic(err)
		}

		rendered := map[string]bool{}
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
					rendered[appInfo.Name] = true
				} else {
					for _, patch := range patchs {
						patchFound = (patch.BashPatch != "" ||
							patch.DockerFromPatch != "" ||
							patch.DockerInstallPatch != "" ||
							patch.GolangBuild != "")
					}
					if len(patchs) == 0 {
						patchFound = len(patchs) == 0
					}
				}
			}
		}

		valuesFiles, err := doublestar.FilepathGlob(fmt.Sprintf("charts/**/values.yaml"))
		if err != nil {
			log.Fatal(err)
		}

		if level >= 0 {
			for i := range valuesFiles {
				path := valuesFiles[i]
				chart := strings.Split(strings.ReplaceAll(path, "\\", "/"), "/")[2]
				if strings.Contains(path, "template") {
					continue
				}

				fileBytes, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}

				fileString := string(fileBytes)

				// version
				repositoryRegex, _ := regexp.Compile("repository:([ \t]*)bitnami/([^ \t\r\n]*)")
				repositorySubmatchGroup := repositoryRegex.FindAllStringSubmatch(fileString, -1)
				dockerFound := true
				for _, groups := range repositorySubmatchGroup {
					if !rendered[groups[2]] {
						dockerFound = false
					}
				}

				if dockerFound {
					emoji.Println(fmt.Sprintf(":heavy_check_mark: %v", chart))
				}
			}
		}

		if level >= 1 {
			for i := range valuesFiles {
				path := valuesFiles[i]
				chart := strings.Split(strings.ReplaceAll(path, "\\", "/"), "/")[2]
				if strings.Contains(path, "template") {
					continue
				}

				fileBytes, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}

				fileString := string(fileBytes)

				// version
				repositoryRegex, _ := regexp.Compile("repository:([ \t]*)bitnami/([^ \t\r\n]*)")
				repositorySubmatchGroup := repositoryRegex.FindAllStringSubmatch(fileString, -1)
				dockerFoundAll := true
				dockerNotFoundAll := len(repositorySubmatchGroup) != 0
				for _, groups := range repositorySubmatchGroup {
					if !rendered[groups[2]] {
						dockerFoundAll = false
					}
					if rendered[groups[2]] {
						dockerNotFoundAll = false
					}
				}

				mainImage := false
				for appName, _ := range rendered {
					if strings.Contains(chart, appName) {
						mainImage = true
					}
					if strings.Contains(appName, chart) {
						mainImage = true
					}
				}

				if !dockerFoundAll && !dockerNotFoundAll && mainImage {
					emoji.Println(fmt.Sprintf(":question: %v", chart))
					for _, groups := range repositorySubmatchGroup {
						if !rendered[groups[2]] {
							emoji.Println(fmt.Sprintf("  :x: %v", groups[2]))
						} else {
							emoji.Println(fmt.Sprintf("  :heavy_check_mark: %v", groups[2]))
						}
					}
				}
			}
		}

		if level >= 2 {
			for i := range valuesFiles {
				path := valuesFiles[i]
				chart := strings.Split(strings.ReplaceAll(path, "\\", "/"), "/")[2]
				if strings.Contains(path, "template") {
					continue
				}

				fileBytes, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}

				fileString := string(fileBytes)

				// version
				repositoryRegex, _ := regexp.Compile("repository:([ \t]*)bitnami/([^ \t\r\n]*)")
				repositorySubmatchGroup := repositoryRegex.FindAllStringSubmatch(fileString, -1)
				dockerFoundAll := true
				dockerNotFoundAll := len(repositorySubmatchGroup) != 0
				for _, groups := range repositorySubmatchGroup {
					if !rendered[groups[2]] {
						dockerFoundAll = false
					}
					if rendered[groups[2]] {
						dockerNotFoundAll = false
					}
				}

				mainImage := false
				for appName, _ := range rendered {
					if strings.Contains(chart, appName) {
						mainImage = true
					}
					if strings.Contains(appName, chart) {
						mainImage = true
					}
				}

				if !dockerFoundAll && !dockerNotFoundAll && !mainImage {
					emoji.Println(fmt.Sprintf(":question: %v", chart))
					for _, groups := range repositorySubmatchGroup {
						if !rendered[groups[2]] {
							emoji.Println(fmt.Sprintf("  :x: %v", groups[2]))
						} else {
							emoji.Println(fmt.Sprintf("  :heavy_check_mark: %v", groups[2]))
						}
					}
				}
			}
		}

		if level >= 3 {
			for i := range valuesFiles {
				path := valuesFiles[i]
				chart := strings.Split(strings.ReplaceAll(path, "\\", "/"), "/")[2]
				if strings.Contains(path, "template") {
					continue
				}

				fileBytes, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}

				fileString := string(fileBytes)

				// version
				repositoryRegex, _ := regexp.Compile("repository:([ \t]*)bitnami/([^ \t\r\n]*)")
				repositorySubmatchGroup := repositoryRegex.FindAllStringSubmatch(fileString, -1)
				dockerFoundAll := true
				dockerNotFoundAll := len(repositorySubmatchGroup) != 0
				for _, groups := range repositorySubmatchGroup {
					if !rendered[groups[2]] {
						dockerFoundAll = false
					}
					if rendered[groups[2]] {
						dockerNotFoundAll = false
					}
				}

				if !dockerFoundAll && dockerNotFoundAll {
					emoji.Println(fmt.Sprintf(":x: %v", chart))
					for _, groups := range repositorySubmatchGroup {
						if !rendered[groups[2]] {
							emoji.Println(fmt.Sprintf("  :x: %v", groups[2]))
						} else {
							emoji.Println(fmt.Sprintf("  :heavy_check_mark: %v", groups[2]))
						}
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkChartsCmd)
	checkChartsCmd.PersistentFlags().IntVar(&level, "level", 0, "level")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkChartsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkChartsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
