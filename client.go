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

	log "github.com/Sirupsen/logrus"
)

const (
	APIVersion = "v2"
)

type Client struct {
	host       *url.URL
	httpClient *http.Client
	username   string
	password   string
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func NewClient(host string, tlsConfig *tls.Config) (*Client, error) {
	h, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("can't parse host %s", host)
	}

	return &Client{
		host:       h,
		httpClient: newHTTPClient(h, tlsConfig),
	}, nil
}

func (c *Client) SetBasicAuth(username, password string) {
	c.username = username
	c.password = password
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

	req, err := http.NewRequest(method, c.host.String()+path, params)
	if err != nil {
		return nil, -1, err
	}

	// Prepare and do the request
	req.Header.Set("User-Agent", "gomarathon")
	req.Header.Set("Content-Type", "application/json")
	if c.username != "" && c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, -1, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}
	if resp.StatusCode >= 400 {
		return nil, resp.StatusCode, fmt.Errorf("status code : %d, body: %s", resp.StatusCode, body)
	}

	return body, resp.StatusCode, nil
}

// request prepare the request by setting the correct methods and parameters
func (c *Client) request(options *RequestOptions) ([]byte, int, error) {
	if options.Method == "" {
		options.Method = "GET"
	}

	path := fmt.Sprintf("/%s/%s", APIVersion, options.Path)

	if options.Params != nil {
		v := url.Values{}

		if options.Params.Host != "" {
			v.Set("host", url.QueryEscape(options.Params.Host))
		}

		if options.Params.Scale {
			v.Set("scale", "true")
		}

		if options.Params.CallbackURL != "" {
			v.Set("callbackUrl", url.QueryEscape(options.Params.CallbackURL))
		}

		if options.Params.Embed != None {
			v.Set("embed", options.Params.Embed.String())
		}

		if options.Params.Cmd != "" {
			v.Set("cmd", url.QueryEscape(options.Params.Cmd))
		}

		if options.Params.Force {
			v.Set("force", "true")
		}

		params := v.Encode()
		if params != "" {
			path = fmt.Sprintf("%s?%s", path, v.Encode())
		}
		log.Debugf("path: %s\n", path)
	}

	return c.do(options.Method, path, options.Datas)
}

func (c *Client) unmarshalJSON(options *RequestOptions, successCodes []int, v interface{}) error {
	data, code, err := c.request(options)
	if err != nil {
		return err
	}
	log.Debugf("response data: %s\n", data)
	if !containsCode(successCodes, code) {
		return fmt.Errorf("status code : %d, data: %s", code, data)
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
		return fmt.Errorf("status code : %d", code)
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
