package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	for {
		reader := make([]byte, 256)
		var writer bytes.Buffer
		nRead, err := conn.Read(reader)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("Error reading connection: ", err.Error())
			os.Exit(1)
		}

		fmt.Println(nRead, string(reader))
		writer.Write([]byte("+PONG\r\n"))
		nWritten, err := conn.Write(writer.Bytes())
		fmt.Println(nWritten, string(writer.String()))
	}
}
