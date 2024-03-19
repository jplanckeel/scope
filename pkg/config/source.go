package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ChartsSource struct {
	Charts           map[string][]string
	ChartsByTagRegex map[string]string `yaml:"charts-by-tag-regex"` // Charts map charts name to regular expression with the charts' tags

}

type configSource map[string]ChartsSource

func NewSource(sourceFile string) (configSource, error) {
	var cfg configSource
	source, err := os.ReadFile(filepath.Clean(sourceFile))
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(source, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to unmarshal %q: %w", sourceFile, err)
	}
	return cfg, nil

}

type chartSearch struct {
	App_version string `yaml:"app_version"`
	Description string `yaml:"description"`
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
}

func chartList(input []byte) ([]chartSearch, error) {
	var y []chartSearch
	err := yaml.Unmarshal(input, &y)
	if err != nil {
		return y, fmt.Errorf("failed to unmarshal %q: %w", input, err)
	}
	return y, nil

}
