package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/genevieve/prise/meeting"
)

const SocketFile = "/tmp/mgr.sock"
const ReportBin = "report"

// The manager sends + receives messages from reports on a unix socket.
func main() {
	if err := os.RemoveAll(SocketFile); err != nil {
		log.Fatal(err)
	}

	cal := new(meeting.Cal)
	if err := rpc.Register(cal); err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	l, err := net.Listen("unix", SocketFile)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println("manager ready")

	if err := http.Serve(l, nil); err != nil {
		log.Fatal(err)
	}
}
