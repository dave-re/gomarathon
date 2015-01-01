package marathon

import (
	"net/http"
)

func (c *Client) GetTasks() (tasks []*Task, err error) {
	options := &RequestOptions{
		Path:   "tasks",
		Method: "GET",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, http.StatusOK, resp)
	tasks = resp.Tasks
	return
}
