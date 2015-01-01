package marathon

import (
	"net/http"
)

func (c *Client) GetInfo() (serverInfo *ServerInfo, err error) {
	options := &RequestOptions{
		Path:   "info",
		Method: "GET",
	}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, serverInfo)
	return
}
