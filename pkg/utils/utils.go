package utils

import "strings"

func EnsureHTTPScheme(url string) string {
	// Check if the URL starts with "http://" or "https://"
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		// If not, prepend "https://"
		url = "https://" + url
	}
	return url
}
