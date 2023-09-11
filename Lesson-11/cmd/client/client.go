package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Printf("Server connection error:%s", err)
		return
	}
	defer conn.Close()

	var message []byte
	var needle string
	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print("Input search string (or 'quit' to exit): ")
		needle, err = reader.ReadString('\n')
		needle = strings.TrimSpace(needle) // Remove leading/trailing whitespace
		fmt.Println("You entered ", needle)

		if err != nil {
			if err == io.EOF {
				fmt.Printf("client closed the connection")
			} else {
				fmt.Printf("server error: %v\n", err)
			}
			break // Exit the loop on error
		}

		if needle == "quit" {
			fmt.Println("Exiting the client.")
			break
		}

		n, err := conn.Write([]byte(needle + "\n"))
		if err != nil {
			fmt.Printf("Error sending request:%s", err)
			continue
		}
		fmt.Printf("%d bytes was sent to server\n", n)

		message, err = serverReader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Error reading server response: %v\n", err)
			break // Exit the loop on error
		}
		if len(message) > 0 {
			fmt.Printf("Search results:%s\n", message)
		}
	}
}
