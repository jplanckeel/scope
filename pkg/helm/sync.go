package helm

import (
	"fmt"
	"os"
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

	// Login to resitry
	if flags.Type != "nexus" {
		err = login(flags)
		if err != nil {
			log.WithField("action", "login").Error(err)
		}

		// Define usernane to namespace
		if flags.Namespace == "" {
			log.WithField("action", "sync").Warnf("setting namespace to %s", flags.Username)
			flags.Namespace = flags.Username
		}
	}

	// Define https scheme if repo do not have scheme
	for repo, charts := range source {
		if !strings.HasPrefix(repo, "https://") {
			repo = "https://" + repo
		}

		for charts, versions := range charts.Charts {
			for _, version := range versions {
				// Check if chart exist in repository source with the version
				if url := findChart(repo, charts, version); url != "" {

					version, err = extractVersion(url)
					if err != nil {
						log.WithField("action", "version").Error(err)
					}

					// Pull chart on source repository
					err := pull(repo, charts, version)
					if err != nil {
						log.WithField("action", "pull").Error(err)
					}

					// Push chart on destination repository
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

					// Delete pulled chart
					err = removeFile(fmt.Sprintf("%s-%s.tgz", charts, version))
					if err != nil {
						log.WithField("action", "removeFile").Error(err)
					}

				}
			}
		}
	}
}

func removeFile(filePath string) (err error) {

	// Check if the file exists
	if _, err = os.Stat(filePath); err == nil {
		// The file exists, so we can delete it
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
		log.Debug("the file has been deleted successfully")
		return

	} else if os.IsNotExist(err) {
		return err
	}
	return fmt.Errorf("error checking the file: %s", err)
}
