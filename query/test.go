package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func main() {

	var doc *Document
	var e error

	if doc, e = NewDocument("http://metalsucks.net"); e != nil {
		log.Fatal(e)
	}

	doc.Find(".reviews-wrap article .review-rhs").Each(func(i int, s *Selection) {
		// For each item found, get the band and title
		band := s.Find("h3").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}
