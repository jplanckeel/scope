package helm

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jplanckeel/scope/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/sourcegraph/run"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

// Function to find chart in repository
func findChart(repository string, chart string, version string) (url string) {

	url, err := repo.FindChartInRepoURL(
		repository,
		chart,
		version,
		"",
		"",
		"",
		getter.All(settings))
	if err != nil {
		log.WithField("action", "findchart").Error(err)
		return
	}
	return
}

// Fucntion to pull chart on repository
func pull(repository string, chart string, version string) error {

	actionConfig := new(action.Configuration)

	registryClient, err := newDefaultRegistryClient(false)
	if err != nil {
		return err
	}
	actionConfig.RegistryClient = registryClient

	client := action.NewPullWithOpts(action.WithConfig(actionConfig))
	client.Version = version
	client.RepoURL = repository
	client.Settings = settings
	if client.Version == "" && client.Devel {
		log.WithField("action", "pull").Warn("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}

	registryClient, err = newRegistryClient(client.CertFile, client.KeyFile, client.CaFile,
		client.InsecureSkipTLSverify, client.PlainHTTP)
	if err != nil {
		return err
	}
	client.SetRegistryClient(registryClient)

	_, err = client.Run(chart)
	if err != nil {
		return err
	}

	log.WithField("action", "pull").Infof("chart %s %s pulled", chart, version)
	return nil
}

// Function to push chart on regsitry
func push(f config.Flags, repoSource string, chart string, version string) error {

	actionConfig := new(action.Configuration)

	registryClient, err := newDefaultRegistryClient(false)
	if err != nil {
		return err
	}
	actionConfig.RegistryClient = registryClient

	client := action.NewPushWithOpts(action.WithPushConfig(actionConfig),
		action.WithTLSClientConfig(f.CertFile, f.KeyFile, f.CaFile),
		action.WithInsecureSkipTLSVerify(f.InsecureSkipTLSverify),
		action.WithPlainHTTP(false))
	client.Settings = settings

	var remote string
	if !f.AppendSource {
		remote = fmt.Sprintf("oci://%s/%s", f.Registry, f.Namespace)
	} else {
		repo, _ := strings.CutPrefix(repoSource, "https://")
		remote = fmt.Sprintf("oci://%s/%s/%s", f.Registry, f.Namespace, repo)
	}

	_, err = client.Run(
		fmt.Sprintf("%s-%s.tgz", chart, version),
		remote,
	)
	if err != nil {
		return err
	}
	log.WithField("action", "push").Infof("chart %s %s pushed", chart, version)
	return nil
}

// Function to push on old regsitry does not support oci format
func pushHttp(f config.Flags, chart string, version string) (err error) {

	var streamLog bytes.Buffer
	err = run.Cmd(
		context.Background(),
		"curl",
		"-T",
		fmt.Sprintf("%s-%s.tgz", chart, version),
		f.Registry,
		"-u",
		fmt.Sprintf("%s:%s", f.Username, f.Password),
	).Run().Stream(&streamLog)
	if err != nil {
		return
	}
	log.WithField("action", "pushHttp").Infof("chart %s %s pushed", chart, version)
	return
}

// Function to login to regsitry
func login(f config.Flags) error {

	actionConfig := new(action.Configuration)

	registryClient, err := newDefaultRegistryClient(false)
	if err != nil {
		return err
	}
	actionConfig.RegistryClient = registryClient

	username, password, err := getUsernamePassword(f.Username, f.Password, f.PasswordFromStdinOpt)
	if err != nil {
		return err
	}

	err = action.NewRegistryLogin(actionConfig).Run(os.Stdout, f.Registry, username, password,
		action.WithCertFile(f.CertFile),
		action.WithKeyFile(f.KeyFile),
		action.WithCAFile(f.CaFile),
		action.WithInsecure(f.InsecureSkipTLSverify))
	if err != nil {
		return err
	}

	return nil
}

func extractVersion(url string) (version string, err error) {
	// Find the index of the last "/" and ".tgz"
	lastSlashIndex := strings.LastIndex(url, "-")
	tgzIndex := strings.Index(url, ".tgz")

	// Extract the necessary parts
	if lastSlashIndex != -1 && tgzIndex != -1 {
		version = url[lastSlashIndex+1 : tgzIndex]
		return
	} else {
		err = fmt.Errorf("invalid string")
		return
	}
}
