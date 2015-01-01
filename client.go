package marathon

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	APIVersion = "v2"
)

type Client struct {
	Host       *url.URL
	HTTPClient *http.Client
}

func NewClient(host string, tlsConfig *tls.Config) (*Client, error) {
	h, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("can't parse host %s", host)
	}

	return &Client{
		Host:       h,
		HTTPClient: newHTTPClient(h, tlsConfig),
	}, nil
}

// do the actual prepared request in request()
func (c *Client) do(method, path string, data interface{}) ([]byte, int, error) {
	var params io.Reader
	var resp *http.Response

	if data != nil {
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, -1, err
		}
		params = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequest(method, c.Host.String()+path, params)
	if err != nil {
		return nil, -1, err
	}

	// Prepare and do the request
	req.Header.Set("User-Agent", "gomarathon")
	req.Header.Set("Content-Type", "application/json")

	resp, err = c.HTTPClient.Do(req)
	if err != nil {
		return nil, -1, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}
	if resp.StatusCode >= 400 {
		return nil, resp.StatusCode, fmt.Errorf("%d: %s", resp.StatusCode, body)
	}

	return body, resp.StatusCode, nil
}

// request prepare the request by setting the correct methods and parameters
func (c *Client) request(options *RequestOptions) ([]byte, int, error) {
	if options.Method == "" {
		options.Method = "GET"
	}

	path := fmt.Sprintf("%s/%s", APIVersion, options.Path)

	if options.Params != nil {
		v := url.Values{}

		if options.Params.Host != "" {
			v.Set("host", url.QueryEscape(options.Params.Host))
		}

		if options.Params.Scale {
			v.Set("scale", "true")
		}

		if options.Params.CallbackURL != "" {
			v.Set("CallbackURL", url.QueryEscape(options.Params.CallbackURL))
		}

		if options.Params.Embed != None {
			v.Set("cmd", url.QueryEscape(options.Params.Cmd))
		}

		if options.Params.Force {
			v.Set("force", "true")
		}

		path = fmt.Sprintf("%s?%s", path, v.Encode())
	}

	return c.do(options.Method, path, options.Datas)
}

func (c *Client) unmarshalJSON(options *RequestOptions, successCodes []int, v interface{}) error {
	data, code, err := c.request(options)
	if err != nil {
		return err
	}
	if !containsCode(successCodes, code) {
		return fmt.Errorf("%d: %s", code, data)
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}

func (c *Client) requestAndCheckSucc(options *RequestOptions, successCodes []int) error {
	_, code, err := c.request(options)
	if err != nil {
		return err
	}
	if !containsCode(successCodes, code) {
		return fmt.Errorf("%d", code)
	}
	return nil
}

func containsCode(successCode []int, code int) bool {
	for _, successCode := range successCode {
		if successCode == code {
			return true
		}
	}
	return false
}
