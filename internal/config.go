package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)


type ChartSyncConfig struct {
	Charts map[string][]string
	ChartsByTagRegex map[string]string      `yaml:"charts-by-tag-regex"` // Charts map charts name to regular expression with the charts' tags
	
}

type sourceConfig map[string]ChartSyncConfig

func newConfig(yamlFile string)(sourceConfig, error) {

	var cfg sourceConfig
	source, err := os.ReadFile(yamlFile)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(source, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to unmarshal %q: %w", yamlFile, err)
	}
	return cfg, nil

}


