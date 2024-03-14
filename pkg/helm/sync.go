package helm

import (
	"strings"

	"github.com/jplanckeel/scope/pkg/config"
	log "github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/cli"
)

// var h *Client
var settings = cli.New()

func Sync(flags config.Flags) {

	//client = NewClient(flags)
	source, err := config.NewSource(flags.SourceFile)
	if err != nil {
		log.Error(err)
	}

	//login to resitry
	login(flags)

	for repo, charts := range source {
		//index := 2
		if !strings.HasPrefix(repo, "https://") {
			repo = "https://" + repo
			//index = 0
		}
		//repoName := strings.Split(repo, "/")

		/*
			err := h.repoAdd(repoName[index], repo)
			if err != nil {
				log.Errorf("can't add repo %v\n", err)
			}*/

		for charts, versions := range charts.Charts {
			for _, version := range versions {
				//chartName := repoName[index] + "/" + chart

				// check if chart exist in repository source with the version
				if url := findChart(repo, charts, version); url != "" {

					version, err = extractVersion(url)
					if err != nil {
						log.WithField("action", "version").Error(err)
					}

					err := pull(repo, charts, version)
					if err != nil {
						log.WithField("action", "pull").Error(err)
					}
					err = push(flags, charts, version)
					if err != nil {
						log.WithField("action", "push").Error(err)
					}
				}
				//pullAndPush(registry, chart, chartName, version)

			}
		}
	}
}

/*
func pullAndPush(registry string, chart string, chartName string, version string) {
	var r string = h.repoDest

	listSource, err := h.searchChart(chartName, version)
	if err != nil {
		log.Errorf("can't search chart in repo %v\n", err)
	}

	chartSource, err := chartList(listSource.Bytes())
	if err != nil {
		log.Errorf("can't list chart in repo %v\n", err)
	}

	for _, c := range chartSource {

		if h.registryType == "oci" {
			r = "oci://" + h.repoDest + "/helm-mirrors/" + registry
		}

		err := h.pullChart(c.Name, c.Version)
		if err != nil {
			log.Errorf("can't pull chart from repo %v\n", err)
		}
		err = h.pushChart(r, chart, c.Version)
		if err != nil {
			log.Errorf("can't push chart to repo %v\n", err)
		}

	}
}
*/
