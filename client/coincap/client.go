package coincap

import (
	"errors"
	"net/http"
	"os"
	"time"
)

type Client struct {
	baseUrl string
	http    *http.Client
}

func (c *Client) GetAsset(id string) (Asset, error) {
	r, err := c.http.Get(c.baseUrl + "/assets/" + id)
	if err != nil {
		return Asset{}, err
	}

	body, err := new(response[Asset]).toJson(*r)
	if err != nil {
		return Asset{}, err
	}

	return body.Data, nil
}

func (c *Client) GetAssets() ([]Asset, error) {
	r, err := c.http.Get(c.baseUrl + "/assets/")
	if err != nil {
		return []Asset{}, err
	}

	body, err := new(response[[]Asset]).toJson(*r)
	if err != nil {
		return []Asset{}, err
	}

	return body.Data, nil
}

func NewClient(timeout time.Duration) (Client, error) {
	if timeout == 0 {
		return Client{}, errors.New("timeout shouldn't be zero")
	}
	return Client{
		baseUrl: "https://api.coincap.io/v2",
		http: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next: &bearerRoundTripper{
					next: http.DefaultTransport,
				},
			},
		},
	}, nil
}
