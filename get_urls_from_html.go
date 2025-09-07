package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string
	h := strings.NewReader(htmlBody)
	doc, err := html.Parse(h)
	if err != nil {
		return nil, err
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					cu, err := cleanUrl(a.Val, rawBaseURL)
					if err != nil {
						return nil, err
					}
					urls = append(urls, cu)
					break
				}
			}
		}
	}

	return urls, nil
}

func cleanUrl(u, baseUrl string) (string, error) {
	pu, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(u, "http") {
		pu, err = url.Parse(baseUrl)
		pu.Path = u

		if err != nil {
			return "", err
		}
	}
	return pu.String(), nil
}
