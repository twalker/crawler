package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := map[string]struct {
		input       string
		expected    string
		expectError bool
	}{
		"trailing slash": {
			input:       "https://blog.boot.dev/path/",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		"valid": {
			input:       "https://blog.boot.dev/path",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},

		"invalid": {
			input:       ":invalid:",
			expected:    "",
			expectError: true,
		},
		"with http": {
			input:       "http://blog.boot.dev/path/",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		"http with trailing": {
			input:       "http://blog.boot.dev/path",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := normalizeURL(tt.input)
			if err != nil && !tt.expectError {
				t.Errorf("Test '%s' FAIL: unexpected error: %v", name, err)
			}
			if actual != tt.expected {
				t.Errorf("expected: %s, but got: %s", tt.expected, actual)
			}
		})
	}
}
