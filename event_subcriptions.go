package gomarathon

import (
	"net/http"
)

// EventResponse represents response for event subscriptions API
type EventResponse struct {
	CallbackURL string `json:"callbackUrl,omitempty"`
	ClientIP    string `json:"clientIp,omitempty"`
	EventType   string `json:"eventType,omitempty"`
}

// RegisterCallbackURL register a callback URL as an event subscriber
// http://goo.gl/pJthfu
func (c *Client) RegisterCallbackURL(callbackURL string) (response *EventResponse, err error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "POST",
		Params: &Parameters{CallbackURL: callbackURL},
	}
	response = &EventResponse{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, response)
	return
}

// GetCallbackURLs gets all event subscriber callback URLs
// http://goo.gl/2ye529
func (c *Client) GetCallbackURLs() (callbackURLs []string, err error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "GET",
	}
	resp := &response{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, resp)
	callbackURLs = resp.CallbackUrls
	return
}

// UnregisterCallbackURL Unregister a callback URL from the event subscribers list
// http://goo.gl/dM6rPB
func (c *Client) UnregisterCallbackURL(callbackURL string) (response *EventResponse, err error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "DELETE",
		Params: &Parameters{CallbackURL: callbackURL},
	}
	response = &EventResponse{}
	err = c.unmarshalJSON(options, []int{http.StatusOK}, response)
	return
}
