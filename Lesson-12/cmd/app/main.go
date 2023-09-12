package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"thinknetica/Lesson-12/pkg/crawler"
	"thinknetica/Lesson-12/pkg/crawler/spider"
	"thinknetica/Lesson-12/pkg/index"
	"thinknetica/Lesson-12/pkg/webapp"
	"time"

	"github.com/gorilla/mux"
)

const depth = 2
const serverAddr = "localhost:8081" // "0.0.0.0:8081"

func main() {
	scanResults, err := performCrawling()
	if err != nil {
		fmt.Printf("Error during crawling: %v", err)
	}

	sortedResults := prepareScanResults(scanResults)

	reverseIndexService := index.NewReverseIndex()
	reverseIndex := reverseIndexService.Build(sortedResults)

	controller := webapp.NewController(reverseIndex, sortedResults)

	err = startHTTPServer(controller)
	if err != nil {
		fmt.Printf("Error starting HTTP server: %v", err)
	}
}

func performCrawling() ([]crawler.Document, error) {
	resources := []string{"https://golang-org.appspot.com/", "https://go.dev/"}
	s := spider.New()
	var scanResults []crawler.Document
	for _, url := range resources {
		result, err := s.Scan(url, depth)
		if err != nil {
			return nil, fmt.Errorf("error scanning docs in %s resource: %v", url, err)
		}
		scanResults = append(scanResults, result...)
	}
	return scanResults, nil
}

func prepareScanResults(scanResults []crawler.Document) []crawler.Document {
	source := rand.NewSource(time.Now().UnixNano())
	randSource := rand.New(source)
	for idx := range scanResults {
		scanResults[idx].ID = randSource.Int()
	}

	sort.Slice(scanResults, func(i, j int) bool {
		return scanResults[i].ID < scanResults[j].ID
	})
	return scanResults
}

func startHTTPServer(controller *webapp.Controller) error {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.ShowIndexData) // Маршрут по умолчанию
	router.HandleFunc("/index", controller.ShowIndexData)
	router.HandleFunc("/docs", controller.ShowDocData)
	router.HandleFunc("/index.xml", controller.ShowIndexDataXML)
	router.HandleFunc("/index.html", controller.ShowIndexDataHTML)
	router.HandleFunc("/docs.xml", controller.ShowDocData)
	router.HandleFunc("/docs.html", controller.ShowDocData)
	fmt.Printf("Starting HTTP server on %s...\n", serverAddr)
	return http.ListenAndServe(serverAddr, router)
}
