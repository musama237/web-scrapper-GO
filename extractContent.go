package main

import "github.com/PuerkitoBio/goquery"

type Book struct {
	Title  string
	Price  string
	Rating string
}

func extractData(doc *goquery.Document) ([]Book, error) {
	var books []Book

	// Find and iterate over each book element
	doc.Find(".product_pod").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3 a").AttrOr("title", "")          // Extract title
		price := s.Find(".price_color").Text()               // Extract price
		rating := s.Find(".star-rating").AttrOr("class", "") // Extract rating class

		// Create a Book struct and append it to the books slice
		books = append(books, Book{
			Title:  title,
			Price:  price,
			Rating: parseRating(rating), // Assuming you have a function to parse rating
		})
	})

	return books, nil
}
