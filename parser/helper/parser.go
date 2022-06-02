package helper

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Parse(url string) ([]string, error) {

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Checking status
	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		return nil, errors.New("status code is not expected")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, errors.New("document body not readable")
	}

	// Find all anchor items
	anchors := doc.Selection.Find("a")

	// Created an array to saving all anchors' text
	var anchorsText []string

	for _, a := range anchors.Nodes {
		anchorsText = append(anchorsText, a.FirstChild.Data)
	}

	return anchorsText, nil
}
