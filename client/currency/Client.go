package currency

import (
	"io"
	"net/http"
	"os"
)

const (
	defaultBaseUrl = "https://api.currencyfreaks.com"
	defaultApiKey  = ""
)

type Client struct {
	baseUrl string
	http    *http.Client
}

func (c *Client) GetLatest() (string, error) {
	r, err := c.http.Get(defaultBaseUrl + "/latest")

	if err != nil {
		return "nil", err
	}

	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)

	if err != nil {
		return "error", err
	}

	return string(b), nil
}

func NewClient() Client {
	return Client{
		baseUrl: defaultBaseUrl,
		http: &http.Client{
			Transport: &authRoundTripper{
				apiKey: defaultApiKey,
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}
}
