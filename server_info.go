package marathon

import (
	"net/http"
)

func (c *Client) GetInfo() (serverInfo *ServerInfo, err error) {
	options := &RequestOptions{
		Path:   "info",
		Method: "GET",
	}
	serverInfo = &ServerInfo{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, serverInfo)
	return
}
