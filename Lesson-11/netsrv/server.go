package netsrv

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"thinknetica/Lesson-11/pkg/crawler"
	"time"
)

type Server struct {
	netListener net.Listener
	scanResults []crawler.Document
}

func NewServer(netListener net.Listener, scanResults []crawler.Document) *Server {
	return &Server{netListener: netListener, scanResults: scanResults}
}

func (s *Server) ListenAndServe() {
	defer s.netListener.Close()
	fmt.Println("Server is ready to accept client connections")
	for {
		conn, err := s.netListener.Accept()
		if err != nil {
			fmt.Printf("Accept connection error: %s\n", err)
			continue
		}

		// Limit the number of concurrent connections (adjust as needed)
		go s.handleClientRequest(conn)
	}
}

func (s *Server) handleClientRequest(connection net.Conn) {
	defer connection.Close()
	var err error
	var needle string
	reader := bufio.NewReader(connection)

	for {
		needle, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading client request: %s\n", err)
			return
		}

		needle = strings.TrimSpace(needle)
		searchResult := s.search(needle)
		if len(searchResult) == 0 {
			result := fmt.Sprintf(" No data found by: %s", needle)
			_, err := connection.Write([]byte(result + "\n"))
			if err != nil {
				fmt.Printf("Error writing response to client: %s\n", err)
			}
			continue
		}

		result := strings.Join(searchResult, ",")
		_, err = connection.Write([]byte(result + "\n"))

		if err != nil {
			fmt.Printf("Error writing response to client: %s\n", err)
		}

		// Reset the connection deadline on each successful request
		connection.SetDeadline(time.Now().Add(time.Second * 60))
	}
}

func (s *Server) search(needle string) []string {
	var links []string
	for _, value := range s.scanResults {
		if strings.Contains(value.Title, needle) || strings.Contains(value.Body, needle) {
			links = append(links, fmt.Sprintf("%s %s", value.Title, value.Body))
		}
	}
	return links
}
