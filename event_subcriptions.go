package marathon

import (
	"net/http"
)

type EventResponse struct {
	CallbackUrl string `json:"callbackUrl,omitempty"`
	ClientIp    string `json:"clientIp,omitempty"`
	EventType   string `json:"eventType,omitempty"`
}

func (c *Client) RegisterCallbackURL(callbackURL string) (response *EventResponse, err error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "POST",
		Params: &Parameters{CallbackURL: callbackURL},
	}
	err = c.unmarshalJSON(options, []int{http.StatusCreated}, response)
	return
}

func (c *Client) GetCallbackURLs() (callbackURLs []string, err error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "GET",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusCreated}, resp)
	callbackURLs = resp.CallbackUrls
	return
}

func (c *Client) UnregisterCallbackURL(callbackURL string) (response *EventResponse, err error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "DELETE",
		Params: &Parameters{CallbackURL: callbackURL},
	}
	err = c.unmarshalJSON(options, []int{http.StatusCreated}, response)
	return
}
