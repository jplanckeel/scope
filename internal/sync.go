package internal

import (
	"fmt"
	"strings"
)

var h *Client

func Sync(binaryHelm string, configPath string, repoDest string, registryType string, user string, password string, dryrun bool) {

	h = NewClient(binaryHelm, repoDest, registryType, user, password, dryrun)
	cfg, err := newConfig(configPath)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	for registry, charts := range cfg {
		repo := registry
		if !strings.HasPrefix(repo, "https://") {
			repo = "https://" + repo
		}
		repoName := strings.Split(registry, "/")
		err := h.repoAdd(repoName[0], repo)
		if err != nil {
			fmt.Printf("error: can't add repo %v\n", err)
		}

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
		fmt.Printf("error: can't search chart in repo %v\n", err)
	}

	chartSource, err := chartList(listSource.Bytes())
	if err != nil {
		fmt.Printf("error: can't list chart in repo %v\n", err)
	}

	for _, c := range chartSource {

		r := registryD(registry, h.repoDest, h.registryType)

		err := h.pullChart(c.Name, c.Version)
		if err != nil {
			fmt.Printf("error: can't pull chart from repo %v\n", err)
		}
		err = h.pushChart(r, chart, c.Version)
		if err != nil {
			fmt.Printf("error: can't push chart to repo %v\n", err)
		}

	}
}

func registryD(r string, repod string, t string) string {

	if t == "oci" {
		oci := "oci://" + h.repoDest + "/helm-mirrors/" + r
		return oci
	}
	return repod
}
