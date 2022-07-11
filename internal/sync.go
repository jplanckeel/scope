package internal

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sourcegraph/run"
)

func Sync(binaryHelm string, configPath string, repoDest string) {

	h := binaryHelm
	ctx := context.Background()

	cfg, err := newConfig(configPath)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	for registry, charts := range cfg {
		repo := "https://" + registry
		repoName := strings.Split(registry, "/")
		err := run.Cmd(ctx, h, "repo", "add", repoName[0], repo).Run().Stream(os.Stdout)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}

		for chart, versions := range charts.Charts {
			for _, version := range versions {
				chartName := repoName[0] + "/" + chart
				registryDest := "oci://" + repoDest + "/helm-mirrors/" + registry + "/" + chart
				
				err := run.Cmd(ctx, h, "pull", chartName, "--version", version).Run().Stream(os.Stdout)
				if err != nil {
					fmt.Printf("error: %s\n", err.Error())
				}
				chartNameDest := chartName + ".tgz"
				err2 := run.Cmd(ctx, h, "push", chartNameDest, registryDest).Run().Stream(os.Stdout)
				if err2 != nil {
					fmt.Printf("error: %s\n", err.Error())
				}
			}
		}
	}
}
