package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"order-api/config"
	"time"
)

// Criar cliente http
const timeout = 30

type Client struct {
	apiURL *url.URL
	client *http.Client
}

func NewUserClient(cfg config.Config) *Client {
	apiURL, _ := url.Parse(cfg.User.APIURL)
	return &Client{
		apiURL: apiURL,
		client: &http.Client{
			Timeout: timeout * time.Second,
		},
	}
}

func (c *Client) NewRequest(ctx context.Context, method, url string, body interface{}) (*http.Request, error) {
	const op = "Client.User.NewUserClient"

	u, err := c.apiURL.Parse(url)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return c.buildRequest(ctx, method, u, buf)
}

func (c *Client) buildRequest(ctx context.Context, method string, u *url.URL, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf(`error creating request: %w`, err)
	}

	return req, nil
}

func (c *Client) DoRequest(req *http.Request) (*http.Response, error) {
	const op = "Client.User.DoRequest"

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf("%s: %w", op, fmt.Errorf("unexpected status code: %d", res.StatusCode))
	}

	return res, nil
}
