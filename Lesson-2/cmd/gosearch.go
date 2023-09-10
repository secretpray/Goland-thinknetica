package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"Lesson-2/pkg/crawler"

	"Lesson-2/pkg/crawler/spider"
)

var urls = []string{"https://golang.org", "https://go.dev"}

func main() {
	query := flag.String("s", "", "Search")
	flag.Parse()

	if *query == "" {
		fmt.Println("Need use with flag `s` for search. For example: gosearch -s go")
		return
	}

	docs, err := scan(urls)
	if err != nil {
		log.Println(err)
		return
	}

	for _, d := range docs {
		if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*query)) {
			fmt.Printf("`%s` found in url: %s\n", *query, d.URL)
		}
	}
}

func scan(urls []string) ([]crawler.Document, error) {
	var result []crawler.Document

	s := spider.New()
	for _, url := range urls {
		docs, err := s.Scan(url, 2)
		if err != nil {
			log.Println(err)
			return result, err
		}

		for _, i := range docs {
			result = append(result, i)
		}
	}
	return result, nil
}
