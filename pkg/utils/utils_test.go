package utils

import (
	"fmt"
	"testing"
)

func TestEnsureHTTPScheme(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"example.com", "https://example.com"},
		{"http://example.com", "http://example.com"},
		{"https://example.com", "https://example.com"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Testing URL: %s", tc.input), func(t *testing.T) {
			result := EnsureHTTPScheme(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
		})
	}
}
