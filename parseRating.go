package main

import "strings"

func parseRating(ratingClass string) string {
	// A simple switch case to parse rating from class name
	switch {
	case strings.Contains(ratingClass, "One"):
		return "1"
	case strings.Contains(ratingClass, "Two"):
		return "2"
	case strings.Contains(ratingClass, "Three"):
		return "3"
	case strings.Contains(ratingClass, "Four"):
		return "4"
	case strings.Contains(ratingClass, "Five"):
		return "5"
	default:
		return "0"
	}
}
