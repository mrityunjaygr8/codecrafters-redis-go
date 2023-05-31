package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/mrityunjaygr8/app/proto"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	log.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConn(conn)

	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		// value, err := proto.DecodeRESP(bufio.NewReader(conn))
		value, err := proto.DecodeRESP(bufio.NewReader(conn))
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			fmt.Println("Error decoding RESP: ", err.Error())
			return
		}

		command := value.Array()[0].String()
		args := value.Array()[1:]

		switch command {
		case "ping":
			conn.Write([]byte("+PONG\r\n"))
		case "echo":
			conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(args[0].String()), args[0].String())))
		default:
			conn.Write([]byte("-ERR wierd shit '" + command + "'\r\n"))
		}
	}

}
