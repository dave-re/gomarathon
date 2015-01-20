package gomarathon

import (
	"net/http"
)

// GetTaskQueue show content of the task queue
// http://goo.gl/6TULb6
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
