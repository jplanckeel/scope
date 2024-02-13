package internal

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

var h *Client

func Sync(config ScopeConfig) {

	h = NewClient(config)
	cfg, err := newSourceConfig(config.ConfigFile)
	if err != nil {
		log.Error(err)
	}

	for registry, charts := range cfg {
		index := 2
		repo := registry
		if !strings.HasPrefix(repo, "https://") {
			repo = "https://" + repo
			index = 0
		}
		repoName := strings.Split(registry, "/")
		err := h.repoAdd(repoName[index], repo)
		if err != nil {
			log.Errorf("can't add repo %v\n", err)
		}

		for chart, versions := range charts.Charts {
			for _, version := range versions {
				chartName := repoName[index] + "/" + chart

				pullAndPush(registry, chart, chartName, version)

			}
		}
	}
}

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
