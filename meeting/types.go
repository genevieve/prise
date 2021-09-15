package meeting

import (
	"errors"
	"fmt"
)

type Cal struct {
	Meetings      map[int]Meeting
	NextMeetingID int
}

func NewCal() *Cal {
	return &Cal{
		Meetings:      make(map[int]Meeting),
		NextMeetingID: 1,
	}
}

type Meeting struct {
	ID    int
	Notes string
	Date  string
}

type Request struct {
	Date string
}

type Response struct {
	MeetingID int
	Date      string
}

func (c *Cal) Meeting(req *Request, resp *Response) error {
	if req.Date == "" {
		return errors.New("invalid meeting req date")
	}
	m := Meeting{}
	m.ID = c.NextMeetingID
	c.NextMeetingID++
	m.Date = req.Date
	c.Meetings[m.ID] = m

	resp.MeetingID = m.ID
	resp.Date = m.Date

	return nil
}

type NotesRequest struct {
	MeetingID int
	Notes     string
}

type NotesResponse struct {
	Ack bool
}

func (c *Cal) MeetingNotes(req *NotesRequest, resp *NotesResponse) error {
	resp.Ack = true
	if m, ok := c.Meetings[req.MeetingID]; ok {
		m.Notes = req.Notes
		c.Meetings[req.MeetingID] = m
	} else {
		return fmt.Errorf("meeting id %d not found", req.MeetingID)
	}
	return nil
}
