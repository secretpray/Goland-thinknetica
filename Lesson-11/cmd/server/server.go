package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"thinknetica/Lesson-11/netsrv"
	"thinknetica/Lesson-11/pkg/crawler"
	"thinknetica/Lesson-11/pkg/crawler/spider"
)

const (
	depth         = 2
	serverAddress = "0.0.0.0:8000"
)

type ConnectionHandler func(net.Conn, []crawler.Document)

func main() {
	resources := []string{"https://golang-org.appspot.com/", "https://go.dev/"}
	s := spider.New()
	var scanResults []crawler.Document

	fmt.Println("Start scanning resorces")
	for _, url := range resources {
		result, err := s.Scan(url, depth)
		if err != nil {
			fmt.Printf("Error due to scanning docs in %s resourse: %s", url, err)
			continue
		}
		scanResults = append(scanResults, result...)
	}

	fmt.Println("Scanning is finished")
	fmt.Println("Starting tcp server")

	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Listener error:%s", err)
		return
	}

	netListener := netsrv.NewServer(listener, scanResults)
	go netListener.ListenAndServe() // or go netListener.ListenAndServe()

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh // Wait for a signal
	fmt.Println("Shutting down the server...")

	listener.Close() // Close any active connections

	fmt.Println("Server is gracefully stopped")
}
