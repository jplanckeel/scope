package helm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractVersion(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected string
		wantErr  bool
	}{
		{"valid archive", "https://example.com/charts/prometheus-1.0.0.tgz", "1.0.0", false},
		{"valid archive with path", "oci://repo/charts/mychart-2.3.4.tgz", "2.3.4", false},
		{"missing tgz", "https://example.com/charts/prometheus-1.0.0", "", true},
		{"missing dash", "https://example.com/charts/prometheus.tgz", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			version, err := extractVersion(tc.url)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Empty(t, version)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, version)
			}
		})
	}
}
