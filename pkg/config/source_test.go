package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestNewSource(t *testing.T) {
	goodConfig := `apache.github.io/superset:
  charts:
    superset:
      - 0.1.0
      - 0.1.1
prometheus-community.github.io/helm-charts:
  charts:
    prometheus:
      - ~11.1.0
    prometheus-node-exporter:
      - 2.0.0
      - 2.0.1
`

	tmpDir := t.TempDir()
	sourceFile := filepath.Join(tmpDir, "good.yaml")
	assert.NoError(t, os.WriteFile(sourceFile, []byte(goodConfig), 0o644))

	expected := configSource{
		"apache.github.io/superset": {
			Charts: map[string][]string{
				"superset": {"0.1.0", "0.1.1"},
			},
		},
		"prometheus-community.github.io/helm-charts": {
			Charts: map[string][]string{
				"prometheus":               {"~11.1.0"},
				"prometheus-node-exporter": {"2.0.0", "2.0.1"},
			},
		},
	}

	cfg, err := NewSource(sourceFile)
	assert.NoError(t, err)
	assert.Equal(t, expected, cfg)
}

func TestNewSourceFileNotFound(t *testing.T) {
	_, err := NewSource("does-not-exist.yaml")
	assert.Error(t, err)
}

func TestNewSourceInvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	sourceFile := filepath.Join(tmpDir, "bad.yaml")
	assert.NoError(t, os.WriteFile(sourceFile, []byte("not: [ valid: yaml"), 0o644))

	_, err := NewSource(sourceFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal")
}

func TestChartList(t *testing.T) {
	testCase := []chartSearch{
		{App_version: "1.0.0", Description: "chart description", Name: "prometheus", Version: "1.0.0"},
		{App_version: "1.2.0", Description: "chart description", Name: "prometheus", Version: "1.0.0"},
	}

	y, err := yaml.Marshal(testCase)
	assert.NoError(t, err)

	list, err := chartList(y)
	assert.NoError(t, err)
	assert.Equal(t, testCase, list)
}

func TestChartListInvalidYAML(t *testing.T) {
	_, err := chartList([]byte("not: [ valid: yaml"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal")
}
