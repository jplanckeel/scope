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
	log.WithField("action", "repoAdd").Infof("add repo name %s repo %s\n", repoName, repo)
	var streamLog bytes.Buffer
	err := run.Cmd(c.ctx, c.binary, "repo", "add", repoName, repo).Run().Stream(&streamLog)
	if err != nil {
		return err

	}
	log.WithField("action", "repoAdd").Infof("%s", streamLog.String())
	return nil
}

func (c Client) searchChart(chart string, version string) (bytes.Buffer, error) {
	log.WithField("action", "searchChart").Infof("search chart %s version %s\n", chart, version)
	var streamData bytes.Buffer
	err := run.Cmd(c.ctx, c.binary, "search", "repo", chart, "--version", version, "-o", "yaml").Run().Stream(&streamData)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return streamData, nil
}

func (c Client) pullChart(chartName string, version string) error {
	var streamLog bytes.Buffer
	log.WithField("action", "pullChart").Infof("pull chart %s:%s\n", chartName, version)
	if !c.dryrun {
		err := run.Cmd(c.ctx, c.binary, "pull", chartName, "--version", version).Run().Stream(&streamLog)
		if err != nil {
			log.WithField("action", "pullChart").Error(streamLog.String())
			return err
		}

	}
	return nil
}

func (c Client) pushChart(registry string, chartName string, version string) error {
	var streamLog bytes.Buffer
	log.WithField("action", "pushChart").Infof("push chart %s/%s:%s\n", registry, chartName, version)
	if !c.dryrun {
		chartPackage := chartName + "-" + version + ".tgz"
		if c.registryType == "nexus" {
			auth := c.user + ":" + c.password
			err := run.Cmd(c.ctx, "curl", "-T", chartPackage, registry, "-u", auth, "-s").Run().Stream(&streamLog)
			if err != nil {

				return err
			}
			log.WithField("action", "pushChart").Info(streamLog.String())
		} else {
			err := run.Cmd(c.ctx, c.binary, "push", chartPackage, registry).Run().Stream(&streamLog)
			if err != nil {

				return err
			}
			log.WithField("action", "pushChart").Info(streamLog.String())
		}
		log.WithField("action", "pushChart").Infof("chart %s pushed\n", chartPackage)
		errRemove := os.Remove(chartPackage)
		if errRemove != nil {

			return errRemove
		}
	}
	return nil
}
