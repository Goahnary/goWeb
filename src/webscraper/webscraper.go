package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
);

func main () {
	fmt.Printf("What website should I scrape? ");
	var siteUrl string
	fmt.Scan(&siteUrl)

	fmt.Printf("\n");
	fmt.Printf("What is the element ID you want me to fetch? ");
	var elementId string
	fmt.Scan(&elementId)

	fmt.Printf("\n");
	fmt.Printf("Site: %v", siteUrl);
	fmt.Printf("\n");
	fmt.Printf("Element ID: %v", elementId);
	fmt.Printf("\n");

	var prepend string
	if !strings.Contains(siteUrl, "http") {
		prepend = "http://"
	}

	res, err := http.Get(prepend + siteUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v", doc.Find("#" + elementId).Text());
	fmt.Printf("\n");
}
