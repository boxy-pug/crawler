package main

import (
	"testing"
)

func TestNormalizeUrl(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		baseURL     string
		expected    string
		expectedErr bool
	}{
		{
			name:        "remove scheme",
			inputURL:    "https://blog.boot.dev/path",
			baseURL:     "https://blog.boot.dev",
			expected:    "blog.boot.dev/path",
			expectedErr: false,
		},
		{
			name:        "remove trailing slash",
			inputURL:    "https://blog.boot.dev/path/",
			baseURL:     "https://blog.boot.dev",
			expected:    "blog.boot.dev/path",
			expectedErr: false,
		},
		{
			name:        "not a valid url",
			inputURL:    "lkjdsi7y8weohsd",
			baseURL:     "",
			expected:    "/lkjdsi7y8weohsd",
			expectedErr: false,
		},
		{
			name:        "http and trailing /'s",
			inputURL:    "http://blog.boot.dev/path////",
			baseURL:     "https://blog.boot.dev",
			expected:    "blog.boot.dev/path",
			expectedErr: false,
		},
		{
			name:        "long path, only path",
			inputURL:    "/path/yes/sir/hello",
			baseURL:     "https://blog.boot.dev",
			expected:    "blog.boot.dev/path/yes/sir/hello",
			expectedErr: false,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.baseURL, tc.inputURL)
			if tc.expectedErr {
				if err == nil {
					t.Errorf("Test %v - '%s' FAIL: expected error but got none", i, tc.name)
				}
			}
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)

			}
		})
	}
}
