package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		response := "+PONG\r\n"

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			os.Exit(1)
		}

		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error response client: ", err.Error())
			os.Exit(1)
		}
	}
}
