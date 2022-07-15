package internal

import (
	"fmt"
	"strings"
)

func Sync(binaryHelm string, configPath string, repoDest string, dryrun bool) {

	h := NewClient(binaryHelm, dryrun)
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

				listSource, err := h.searchChart(chartName, version)
				if err != nil {
					fmt.Printf("error - can search chart in repo err:%s", err)
				}

				chartSource, err := chartList(listSource.Bytes())
				if err != nil {
					fmt.Printf("error - %s", err)
				}

				for _, c := range chartSource {
					h.pullChart(c.Name, c.Version)
					if err != nil {
						fmt.Printf("error - %s", err)
					}
					h.pushChart(registry, chart, c.Version, repoDest)
					if err != nil {
						fmt.Printf("error - %s", err)
					}
					
				}
				
	

			}
		}
	}
}
