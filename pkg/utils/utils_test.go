package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureHTTPScheme(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"no scheme", "example.com", "https://example.com"},
		{"http scheme", "http://example.com", "http://example.com"},
		{"https scheme", "https://example.com", "https://example.com"},
		{"empty string", "", "https://"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, EnsureHTTPScheme(tc.input))
		})
	}
}
