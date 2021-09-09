package main

import (
	"fmt"
	"log"
	"net"
)

const SocketFile = "/tmp/echo.sock"

func main() {
	message := []byte("hello, world")

	c, err := net.Dial("unix", SocketFile)
	if err != nil {
		log.Fatalf("dial error: %s", err)
	}
	defer c.Close()
	count, err := c.Write(message)
	if err != nil {
		log.Fatalf("write error: %s", err)
	}
	fmt.Printf("wrote %d bytes\n", count)
}
