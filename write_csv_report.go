package main

import (
	"encoding/csv"
	"os"
	"strings"
)

type row struct {
	page_url, h1, first_paragraph, outgoing_link_urls, image_urls string
}

func writeCSVRport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	csvw := csv.NewWriter(file)
	// write header
	csvw.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"})

	for _, pd := range pages {
		//fmt.Println("Writing record for:", k)
		csvw.Write([]string{pd.URL, pd.H1, pd.FirstParagraph, strings.Join(pd.OutgoingLinks, ";"), strings.Join(pd.ImageURLs, ";")})
	}
	return nil
}
