package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name        string
		expected    []string
		expectError bool
		inputURL    string
		inputBody   string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
<body>
<a href="/path/one">ee
	<span>Boot.dev</span>
</a>
<a href="https://other.com/path/one">
	<span>Boot.dev</span>
</a>
</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		}, {
			name:     "nested URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
<body>
<p>
	<a href="/a">ee
		<span>Boot.dev</span>
	</a>
</p>
<a href="https://other.com/path/b">
	<span>Boot.dev</span>
</a>
<article><section> <h3>
  <a href="https://other.com/path/c">c</a>
</h3> </section></article>
</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/a", "https://other.com/path/b", "https://other.com/path/c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tt.inputBody, tt.inputURL)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected: %s, but got: %s", tt.expected, actual)
			}
		})
	}
}
