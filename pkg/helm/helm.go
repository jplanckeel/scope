package helm

import (
	"fmt"
	"os"
	"strings"

	"github.com/jplanckeel/scope/pkg/config"
	log "github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

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

func push(f config.Flags, chart string, version string) error {

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

	// define usernane to namespace
	if f.Namespace == "" {
		f.Namespace = f.Username
	}
	_, err = client.Run(
		fmt.Sprintf("%s-%s.tgz", chart, version),
		fmt.Sprintf("oci://%s/%s", f.Registry, f.Namespace),
	)
	if err != nil {
		return err
	}
	log.WithField("action", "pushed").Infof("chart %s %s pushed", chart, version)
	return nil
}

func pushHttp(f config.Flags, chart string, version string) error {

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
	_, err = client.Run(fmt.Sprintf("%s-%s.tgz", chart, version), fmt.Sprintf("oci://%s/", f.Registry))
	if err != nil {
		return err
	}
	return nil
}

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

	action.NewRegistryLogin(actionConfig).Run(os.Stdout, f.Registry, username, password,
		action.WithCertFile(f.CertFile),
		action.WithKeyFile(f.KeyFile),
		action.WithCAFile(f.CaFile),
		action.WithInsecure(f.InsecureSkipTLSverify))

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
		err = fmt.Errorf("Invalid string")
		return
	}
}
