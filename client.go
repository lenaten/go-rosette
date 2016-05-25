package rosette

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client implements a AlchemyAPI client.
type Client struct {
	*Config
}

// New client.
func New(config *Config) *Client {
	c := &Client{Config: config}
	return c
}

// call rpc style endpoint.
func (c *Client) call(path string, in interface{}) (io.ReadCloser, error) {
	url := "https://api.rosette.com:443/rest/v1/" + path

	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-RosetteAPI-Key", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	r, _, err := c.do(req)
	return r, err
}

// perform the request.
func (c *Client) do(req *http.Request) (io.ReadCloser, int64, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	if res.StatusCode < 400 {
		return res.Body, res.ContentLength, err
	}

	defer res.Body.Close()

	e := &Error{
		Status:     http.StatusText(res.StatusCode),
		StatusCode: res.StatusCode,
	}

	kind := res.Header.Get("Content-Type")

	if strings.Contains(kind, "text/plain") {
		if b, err := ioutil.ReadAll(res.Body); err == nil {
			e.Summary = string(b)
			return nil, 0, e
		}

		return nil, 0, err
	}

	if err := json.NewDecoder(res.Body).Decode(e); err != nil {
		return nil, 0, err
	}

	return nil, 0, e
}
