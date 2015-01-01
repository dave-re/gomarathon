package marathon

import (
	"net/http"
)

func (c *Client) GetTaskQueue() (queue []*TaskQueue, err error) {
	options := &RequestOptions{
		Path:   "queue",
		Method: "GET",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	queue = resp.Queue
	return
}
