package gomarathon

import (
	"net/http"
)

// GetTasks gets tasks of all applications
// http://goo.gl/TJYJmS
func (c *Client) GetTasks() (tasks []*Task, err error) {
	return c.GetTasksWithParams(NoneStatus)
}

// GetTasksWithParams gets tasks of all applications with parameters
// http://goo.gl/TJYJmS
func (c *Client) GetTasksWithParams(status Status) (tasks []*Task, err error) {
	options := &RequestOptions{
		Path:   "tasks",
		Method: "GET",
		Params: &Parameters{
			Status: status,
		},
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	tasks = resp.Tasks
	return
}
