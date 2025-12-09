package client

import (
	"net/http"
)

type DocClient struct {
	Client *Client
}

func (c *DocClient) Get() (map[string]any, error) {
	req, err := c.Client.newRequest(http.MethodGet, "/doc", nil)
	if err != nil {
		return nil, err
	}
	var resp map[string]any
	_, err = c.Client.do(req, &resp)
	return resp, err
}
