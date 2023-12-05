package main

func main() {
	scrapeBooks("http://books.toscrape.com")
}

func scrapeBooks(url string) {

	fetchPage(url)
	// parseHTML()
	// extractBookDetails()
}
