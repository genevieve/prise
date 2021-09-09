package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const SocketFile = "/tmp/echo.sock"

// The manager sends + receives messages from reports on a unix socket.
func main() {
	if err := os.RemoveAll(SocketFile); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", SocketFile)
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Error on accept: %s", err)
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	received := make([]byte, 0)
	for {
		buf := make([]byte, 512)
		count, err := c.Read(buf)
		received = append(received, buf[:count]...)
		if err != nil {
			fmt.Println(string(received))
			if err != io.EOF {
				log.Fatalf("error on read: %s", err)
			}
			break
		}
	}
}
