package meeting

import "errors"

type Cal struct{}

type Request struct {
	Date string
}

type Response struct {
	Date string
}

func (c *Cal) Meeting(req *Request, resp *Response) error {
	if req.Date == "" {
		return errors.New("invalid meeting req date")
	}
	if req.Date == "today" {
		resp.Date = "tomorrow"
	}

	return nil
}
