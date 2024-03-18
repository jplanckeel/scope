package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestNewConfig(t *testing.T) {

	testCase := configSource{
		"apache.github.io/superset": {
			Charts: map[string][]string{
				"superset": {
					"0.1.0",
					"0.1.1",
				},
			},
		},
		"prometheus-community.github.io/helm-charts": {
			Charts: map[string][]string{
				"prometheus": {
					"~11.1.0",
				},
				"prometheus-node-exporter": {
					"2.0.0",
					"2.0.1",
				},
			},
		},
	}

	sourceFile := "../../test/good.yaml"

	cfg, err := NewSource(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, cfg, testCase)
}

func TestNewConfigError(t *testing.T) {

	sourceFile := "../test/bad.yaml"

	_, err := NewSource(sourceFile)
	assert.Error(t, err)
}

func TestChartList(t *testing.T) {

	testCase := []chartSearch{
		{
			App_version: "1.0.0",
			Description: "chart description",
			Name:        "prometheus",
			Version:     "1.0.0",
		}, {
			App_version: "1.2.0",
			Description: "chart description",
			Name:        "prometheus",
			Version:     "1.0.0",
		},
	}

	y, err := yaml.Marshal(testCase)
	if err != nil {
		log.Fatal(err)
	}
	list, err := chartList(y)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, list, testCase)
}
