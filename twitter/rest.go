package twitter

import (
	"net/http"

	"github.com/garyburd/go-oauth"
)

type Client struct {
	consumer   *oauth.Credentials
	access     *oauth.Credentials
	HttpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	client := &Client{}
	if httpClient == nil {
		client.HttpClient = &http.DefaultClient
	} else {
		client.HttpClient = httpClient
	}
	return client
}

func (c *Client) AddCredentials(consumer, access *oauth.Credentials) {
	c.consumer = consumer
	c.access = access
}
