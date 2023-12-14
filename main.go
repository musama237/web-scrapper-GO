// main.go

package main

import (
	"fmt"
	"log"
)

func main() {
	url := "http://books.toscrape.com"

	htmlContent, err := fetchPAGE(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := parseHTML(htmlContent)
	if err != nil {
		log.Fatal(err)
	}

	books, err := extractData(doc)
	if err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		fmt.Printf("Book: %+v\n", book)
	}

}
