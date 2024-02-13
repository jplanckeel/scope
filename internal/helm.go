package internal

import (
	"bytes"
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sourcegraph/run"
)

type Client struct {
	ctx          context.Context
	binary       string
	repoDest     string
	registryType string
	user         string
	password     string
	dryrun       bool
}

func NewClient(config ScopeConfig) *Client {
	ctx := context.Background()
	return &Client{ctx, config.BinaryHelm, config.Registry, config.RegistryType, config.User, config.Password, config.Dryrun}
}

func (c Client) repoAdd(repoName string, repo string) error {
	if c.dryrun {
		log.Infof("dryrun - to add repo name %s repo %s\n", repoName, repo)
	} else {
		err := run.Cmd(c.ctx, c.binary, "repo", "add", repoName, repo).Run().Stream(os.Stdout)
		if err != nil {
			return err
		}

	}
	return nil
}

func (c Client) searchChart(chart string, version string) (bytes.Buffer, error) {
	var streamData bytes.Buffer
	if c.dryrun {
		log.Infof("dryrun - search chart %s version %s\n", chart, version)
	} else {
		err := run.Cmd(c.ctx, c.binary, "search", "repo", chart, "--version", version, "-o", "yaml").Run().Stream(&streamData)
		if err != nil {
			return bytes.Buffer{}, err
		}

	}
	return streamData, nil
}

func (c Client) pullChart(chartName string, version string) error {
	if c.dryrun {
		log.Infof("dryrun - to pull chart %s:%s\n", chartName, version)
	} else {
		err := run.Cmd(c.ctx, c.binary, "pull", chartName, "--version", version).Run().Stream(os.Stdout)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c Client) pushChart(registry string, chartName string, version string) error {
	if c.dryrun {
		log.Infof("dryrun - push chart %s:%s to%s\n", chartName, version, c.repoDest)
	} else {
		chartPackage := chartName + "-" + version + ".tgz"
		if c.registryType == "nexus" {
			auth := c.user + ":" + c.password
			err := run.Cmd(c.ctx, "curl", "-T", chartPackage, registry, "-u", auth, "-s").Run().Stream(os.Stdout)
			if err != nil {

				return err
			}
		} else {
			err := run.Cmd(c.ctx, c.binary, "push", chartPackage, registry).Run().Stream(os.Stdout)
			if err != nil {

				return err
			}
		}
		log.Infof("chart %s pushed\n", chartPackage)
		errRemove := os.Remove(chartPackage)
		if errRemove != nil {

			return errRemove
		}
	}
	return nil
}
