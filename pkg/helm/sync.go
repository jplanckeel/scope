package helm

import (
	"strings"

	"github.com/jplanckeel/scope/pkg/config"
	log "github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/cli"
)

var settings = cli.New()

func Sync(flags config.Flags) {

	source, err := config.NewSource(flags.SourceFile)
	if err != nil {
		log.Error(err)
	}

	//login to resitry
	if flags.Type != "nexus" {
		err = login(flags)
		if err != nil {
			log.WithField("action", "login").Error(err)
		}

		// define usernane to namespace
		if flags.Namespace == "" {
			log.WithField("action", "sync").Warnf("setting namespace to %s", flags.Username)
			flags.Namespace = flags.Username
		}
	}

	//define https scheme if repo do not have scheme
	for repo, charts := range source {
		if !strings.HasPrefix(repo, "https://") {
			repo = "https://" + repo
		}

		for charts, versions := range charts.Charts {
			for _, version := range versions {
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

					if flags.Type != "nexus" {
						err = push(flags, charts, version)
						if err != nil {
							log.WithField("action", "push").Error(err)
						}
					} else {
						err = pushHttp(flags, charts, version)
						if err != nil {
							log.WithField("action", "pushHttp").Error(err)
						}
					}
				}
			}
		}
	}
}
