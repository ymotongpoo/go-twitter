package twitter

import (
	"encoding/json"
	_ "fmt"
	_ "io/ioutil"
	"net/http"
	"net/url"

	"github.com/garyburd/go-oauth/oauth"
)

// Twitter client struct
type Client struct {
	OAuthClient       *oauth.Client
	HttpClient        *http.Client
	accessCredentials *oauth.Credentials
}

type ResourceInfo struct {
	EndPoint       string
	Authentication bool
	HttpMethod     string
	RequiredArgs   []string
	OptionalArgs   []string
}

func NewClient(httpClient *http.Client) *Client {
	client := &Client{}
	client.OAuthClient = &oauth.Client{
		TemporaryCredentialRequestURI: "http://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "http://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "http://api.twitter.com/oauth/access_token",
	}
	if httpClient == nil {
		client.HttpClient = http.DefaultClient
	} else {
		client.HttpClient = httpClient
	}
	return client
}

func (c *Client) makeAPIRequest(ri *ResourceInfo, v *url.Values, result interface{}) error {
	var resp *http.Response
	var err error
	switch ri.HttpMethod {
	case "GET":
		resp, err = c.OAuthClient.Get(c.HttpClient, c.accessCredentials, ri.EndPoint, *v)
	case "POST":
		resp, err = c.OAuthClient.Post(c.HttpClient, c.accessCredentials, ri.EndPoint, *v)
	}
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddCredentials(consumer, access *oauth.Credentials) {
	c.OAuthClient.Credentials = *consumer
	c.accessCredentials = access
}

func (c *Client) MentionsTimeline() ([]*Tweets, error) {
	ri := ResourceInfoMap["status/mentions_timeline"]
	v := &url.Values{}
	result := []*Tweets{}
	err := c.makeAPIRequest(ri, v, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
