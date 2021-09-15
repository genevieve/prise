package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/genevieve/prise/meeting"
)

const SocketFile = "/tmp/mgr.sock"

func main() {
	client, err := rpc.DialHTTP("unix", SocketFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("report connected to manager")

	req := meeting.Request{Date: "today"}
	var resp meeting.Response
	if err := client.Call("Cal.Meeting", &req, &resp); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Meeting set for: ", resp.Date)
}
