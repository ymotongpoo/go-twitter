package twitter

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	_ "io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/garyburd/go-oauth/oauth"
)

var BufferSize = 10000

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

func parseOptionalParams(ri *ResourceInfo, option map[string]string) (*url.Values, error) {
	v := &url.Values{}
	for _, argName := range ri.OptionalArgs {
		if value, exist := option[argName]; exist {
			v.Add(argName, value)
		}
	}
	return v, nil
}

// NewClient returns a Twitter REST API v1.1 client without credentials.
// httpClient will be used to throw HTTP GET/POST request to each APIs.
// (On GAE/Go you may need to specify appengine/urlfetch.Client here)
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

// makeAPIRequest throws GET/POST request with form values v to specific API
// with reffering ResourceInfo, and stores JSON unmarshaled data into result.
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
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

// makeAPIRequest throws GET/POST request with form values v to specific API
// with reffering ResourceInfo, and stores JSON unmarshaled data into result.
func (c *Client) makeStreamAPIRequest(ri *ResourceInfo, v *url.Values, stream interface{}, errch chan<- error) {
	// TODO(ymotongpoo): Check if stream's type is channel first.
	val := reflect.ValueOf(stream)
	typ := val.Type()
	kind := typ.Kind()

	if kind != reflect.Chan {
		errch <- errors.New("Not a channel")
		return
	}

	var resp *http.Response
	var err error
	switch ri.HttpMethod {
	case "GET":
		resp, err = c.OAuthClient.Get(c.HttpClient, c.accessCredentials, ri.EndPoint, *v)
	case "POST":
		resp, err = c.OAuthClient.Post(c.HttpClient, c.accessCredentials, ri.EndPoint, *v)
	}
	if err != nil {
		errch <- err
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	etype := typ.Elem()
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			errch <- err
		}
		
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		element := reflect.Zero(etype)
		err = json.Unmarshal(line, element)
		if err != nil {
			errch <- err
		}
		if !val.TrySend(element) {
			errch <- errors.New("Can't send element")
		}
	}
	return
}

// AddCredentials adds consumer key & consumer secret pair and
// access token & access token secret pair to a Twitter Client c respectively.
func (c *Client) AddCredentials(consumer, access *oauth.Credentials) {
	c.OAuthClient.Credentials = *consumer
	c.accessCredentials = access
}

// MentionsTimeline throws request to statuses/mentions_timeline.json with
// optional parameters 'option'. 'option' should be a map or valiables.
// https://dev.twitter.com/docs/api/1.1/get/statuses/mentions_timeline
func (c *Client) MentionsTimeline(option map[string]string) ([]*Tweets, error) {
	ri := ResourceInfoMap["statuses/mentions_timeline"]
	v, err := parseOptionalParams(ri, option)
	result := []*Tweets{}
	err = c.makeAPIRequest(ri, v, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// https://dev.twitter.com/docs/api/1.1/get/statuses/user_timeline
func (c *Client) UserTimeline(option map[string]string) ([]*Tweets, error) {
	ri := ResourceInfoMap["statuses/user_timeline"]
	v, err := parseOptionalParams(ri, option)
	result := []*Tweets{}
	err = c.makeAPIRequest(ri, v, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// https://dev.twitter.com/docs/api/1.1/get/statuses/home_timeline
func (c *Client) HomeTimeline(option map[string]string) ([]*Tweets, error) {
	ri := ResourceInfoMap["statuses/home_timeline"]
	v, err := parseOptionalParams(ri, option)
	result := []*Tweets{}
	err = c.makeAPIRequest(ri, v, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// https://dev.twitter.com/docs/api/1.1/get/statuses/retweets_of_me
func (c *Client) RetweetsOfMe(option map[string]string) ([]*Tweets, error) {
	ri := ResourceInfoMap["statuses/retweets_of_me"]
	v, err := parseOptionalParams(ri, option)
	result := []*Tweets{}
	err = c.makeAPIRequest(ri, v, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// https://dev.twitter.com/docs/api/1.1/get/statuses/retweets/%3Aid
func (c *Client) Retweets(id int64, option map[string]string) ([]*Tweets, error) {
	ri := *ResourceInfoMap["statuses/retweets/:id"]
	ri.EndPoint = fmt.Sprintf(ri.EndPoint, id)
	v, err := parseOptionalParams(&ri, option)
	result := []*Tweets{}
	err = c.makeAPIRequest(&ri, v, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Aware that return value is a pointer of Tweets, not a slice of that.
// https://dev.twitter.com/docs/api/1.1/get/statuses/show/%3Aid
func (c *Client) Show(id int64, option map[string]string) (*Tweets, error) {
	ri := *ResourceInfoMap["statuses/show/:id"]
	ri.EndPoint = fmt.Sprintf(ri.EndPoint, id)
	v, err := parseOptionalParams(&ri, option)
	result := &Tweets{}
	err = c.makeAPIRequest(&ri, v, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// https://dev.twitter.com/docs/api/1.1/post/statuses/destroy/%3Aid
func (c *Client) Destroy(id int64, option map[string]string) (*Tweets, error) {
	ri := *ResourceInfoMap["statuses/destroy/:id"]
	ri.EndPoint = fmt.Sprintf(ri.EndPoint, id)
	v, err := parseOptionalParams(&ri, option)
	result := &Tweets{}
	err = c.makeAPIRequest(&ri, v, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// https://dev.twitter.com/docs/api/1.1/post/statuses/update
func (c *Client) Update(status string, option map[string]string) (*Tweets, error) {
	ri := ResourceInfoMap["statuses/update"]
	v, err := parseOptionalParams(ri, option)
	v.Add("status", status)

	result := &Tweets{}
	err = c.makeAPIRequest(ri, v, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// https://dev.twitter.com/docs/api/1.1/post/statuses/retweet/%3Aid
func (c *Client) Retweet(id int64, option map[string]string) (*Tweets, error) {
	ri := *ResourceInfoMap["statuses/retweet/:id"]
	ri.EndPoint = fmt.Sprintf(ri.EndPoint, id)
	v, err := parseOptionalParams(&ri, option)

	result := &Tweets{}
	err = c.makeAPIRequest(&ri, v, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateWithMedia posts status with PNG, GIF, JPEG files.
// media is a slice of path to uploading media.
// https://dev.twitter.com/docs/api/1.1/post/statuses/update_with_media
//
// TODO(ymotongpoo): Add file existence validation.
func (c *Client) UpdateWithMedia(status string, media []string, option map[string]string) (*Tweets, error) {
	ri := ResourceInfoMap["statuses/update_with_media"]
	v, err := parseOptionalParams(ri, option)
	v.Add("status", status)
	for _, m := range media {
		v.Add("media[]", m)
	}

	result := &Tweets{}
	err = c.makeAPIRequest(ri, v, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// https://dev.twitter.com/docs/api/1.1/get/statuses/oembed
func (c *Client) OEmbed(params map[string]string) (*OEmbed, error) {
	ri := ResourceInfoMap["statuses/oembed"]
	v, err := parseOptionalParams(ri, params)
	for _, arg := range ri.RequiredArgs {
		if value, exist := params[arg]; exist {
			v.Add(arg, value)
		}
	}
	result := &OEmbed{}
	err = c.makeAPIRequest(ri, v, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// https://dev.twitter.com/docs/api/1.1/get/search/tweets
func (c *Client) Tweets(q string, option map[string]string) (*Search, error) {
	ri := ResourceInfoMap["search/tweets"]
	v, err := parseOptionalParams(ri, option)
	v.Add("q", q)

	result := &Search{}
	err = c.makeAPIRequest(ri, v, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Filter returns a channel of *Tweets which matches one or more filter predicates
// specified in option. At least one predicate parameter ("follow", "locations" or "track")
// must be specified.
//
// https://dev.twitter.com/docs/api/1.1/post/statuses/filter
func (c *Client) Filter(option map[string]string) (<-chan *Tweets, <-chan error) {
	ri := ResourceInfoMap["statuses/filter"]
	v, err := parseOptionalParams(ri, option)
	if err != nil {
		return nil, nil
	}
	// TODO(ymotongpoo): Add option validator to check if a required one exists in the option.

	stream := make(chan *Tweets, BufferSize)
	errch := make(chan error, BufferSize)
	// TODO(ymotongpoo): Implement go function call for stream.
	go c.makeStreamAPIRequest(ri, v, stream, errch)
	return stream, errch
}
