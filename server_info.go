package gomarathon

import (
	"net/http"
)

// GetInfo get info about the Marathon Instance
// http://goo.gl/JWnISL
func (c *Client) GetInfo() (serverInfo *ServerInfo, err error) {
	options := &RequestOptions{
		Path:   "info",
		Method: "GET",
	}
	serverInfo = &ServerInfo{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, serverInfo)
	return
}
