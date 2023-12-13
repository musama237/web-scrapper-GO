package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseHTML(htmlContent string) (*goquery.Document, error) {

	reader := strings.NewReader(htmlContent)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
