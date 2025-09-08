package main

import (
	"net/url"
	"strings"
)

func normalizeURL(u string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	parsed.Path = strings.TrimSuffix(parsed.Path, "/")
	// return fmt.Sprintf("%s%s", parsed.Host, parsed.Path), nil
	return parsed.String(), nil
}
