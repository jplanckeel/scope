package internal

import (
	"fmt"
	"strings"
)

var h *Client

func Sync(binaryHelm string, configPath string, repoDest string, dryrun bool) {

	h = NewClient(binaryHelm, repoDest, dryrun)
	cfg, err := newConfig(configPath)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	for registry, charts := range cfg {
		repo := "https://" + registry
		repoName := strings.Split(registry, "/")
		h.repoAdd(repoName[0], repo)

		for chart, versions := range charts.Charts {
			for _, version := range versions {
				chartName := repoName[0] + "/" + chart

				pullAndPush(registry, chart, chartName, version)

			}
		}
	}
}

func pullAndPush(registry string, chart string, chartName string, version string) {

	listSource, err := h.searchChart(chartName, version)
	if err != nil {
		fmt.Printf("error - can search chart in repo err:%s", err)
	}

	chartSource, err := chartList(listSource.Bytes())
	if err != nil {
		fmt.Printf("error - %s", err)
	}

	for _, c := range chartSource {

		oci := "oci://" + h.repoDest + "/helm-mirrors/" + registry

		h.pullChart(c.Name, c.Version)
		if err != nil {
			fmt.Printf("error - %s", err)
		}
		h.pushChart(oci, chart, c.Version)
		if err != nil {
			fmt.Printf("error - %s", err)
		}

	}
}
