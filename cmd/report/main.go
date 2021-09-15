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

	fmt.Printf("meeting set for: %s\n", resp.Date)

	notesReq := meeting.NotesRequest{
		MeetingID: resp.MeetingID,
		Notes:     "something important",
	}
	var notesResp meeting.NotesResponse
	if err := client.Call("Cal.MeetingNotes", &notesReq, &notesResp); err != nil {
		log.Fatal(err)
	}

	fmt.Println("meeting notes accepted")
}
