package main

import (
	"io/ioutil"
	"log"

	"github.com/hoisie/mustache"
	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Name string   `yaml:"name"`
	On   []string `yaml:"on"`
	Env  struct {
		REGISTRY string `yaml:"REGISTRY"`
	} `yaml:"env"`
	Jobs struct {
		BuildBitnamiCompat struct {
			RunsOn      string `yaml:"runs-on"`
			If          string `yaml:"if"`
			Permissions struct {
				Contents string `yaml:"contents"`
				Packages string `yaml:"packages"`
			} `yaml:"permissions"`
			Strategy struct {
				Matrix struct {
					Include []struct {
						APP          string `yaml:"APP"`
						OSFLAVOUR    string `yaml:"OS_FLAVOUR"`
						VERSIONMAJOR string `yaml:"VERSION_MAJOR"`
						VERSION      string `yaml:"VERSION"`
						REVISION     string `yaml:"REVISION"`
						OSNAME       string `yaml:"OS_NAME"`
						REPOOWNER    string `yaml:"REPO_OWNER"`
					} `yaml:"include"`
				} `yaml:"matrix"`
			} `yaml:"strategy"`
			Steps []struct {
				Name  string `yaml:"name"`
				Uses  string `yaml:"uses,omitempty"`
				Shell string `yaml:"shell,omitempty"`
				Run   string `yaml:"run,omitempty"`
				ID    string `yaml:"id,omitempty"`
				With  struct {
				} `yaml:"with,omitempty"`
			} `yaml:"steps"`
		} `yaml:"build-bitnami-compat"`
	} `yaml:"jobs"`
}

func main() {
	buf, err := ioutil.ReadFile(".github/workflows/build-on-push.yml")
	if err != nil {
		log.Fatal(err)
	}

	workflow := Workflow{}
	err = yaml.Unmarshal(buf, &workflow)

	for _, w := range workflow.Jobs.BuildBitnamiCompat.Strategy.Matrix.Include {
		data := mustache.Render("* {{APP}} : [`{{VERSION_MAJOR}}`, `{{VERSION_MAJOR}}-{{OS_FLAVOUR}}`, `{{VERSION}}`, `{{VERSION}}-{{OS_FLAVOUR}}`, `{{VERSION}}-{{OS_FLAVOUR}}-r{{REVISION}}`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2F{{APP}})",
			map[string]string{
				"APP":           w.APP,
				"OS_FLAVOUR":    w.OSFLAVOUR,
				"VERSION_MAJOR": w.VERSIONMAJOR,
				"VERSION":       w.VERSION,
				"REVISION":      w.REVISION,
				"OS_NAME":       w.OSNAME,
				"REPO_OWNER":    w.REPOOWNER,
			})
		println(data)
	}
}
