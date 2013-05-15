package octogopher

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const defaultApiRoot = "https://api.github.com"
const defaultUserAgent = "OCTOGOPHER"

func NewClient() *Client {
	return &Client{
		ApiRoot:   defaultApiRoot,
		UserAgent: defaultUserAgent,
	}
}

func NewClientBasicAuth(login, password string) *Client {
	c := NewClient()
	c.login = login
	c.password = password
	return c
}

// TODO: Is login necessary for oauth authorization?
func NewClientOAuth(token string) *Client {
	c := NewClient()
	c.oauthToken = token
	return c
}

// TODO: Deal with null json fields as required by api.

type Client struct {
	ApiRoot   string
	UserAgent string

	login      string
	password   string
	oauthToken string
}

// TODO: Use whichever auth is enabled
func (c *Client) authenticateRequest(req *http.Request) {
	if c.password != "" {
		req.SetBasicAuth(c.login, c.password)
	}
}

func (c *Client) NewRequest(method, relativeURL string, body io.Reader) (*http.Request,
	error) {

	fullURL := c.ApiRoot + relativeURL

	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, err
	}

	c.authenticateRequest(req)

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", c.UserAgent)

	// Make sure content length is set
	if req.Body != nil && req.ContentLength <= 0 {
		content, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.ContentLength = int64(len(content))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// TODO: Should params have a different type?
func (c *Client) Get(relativeURL string, params url.Values) (*http.Response, error) {
	pathAndParams := strings.Join([]string{relativeURL, params.Encode()}, "?")

	req, err := c.NewRequest("GET", pathAndParams, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	// TODO: Should this be in the Client.Do() method?
	// TODO: Should http errors be returned as errors?
	// TODO: clean up status code catching stuff
	// TODO: catch 500 errors
	if resp.StatusCode%400 < 100 {
		return nil, ClientError(resp)
	}

	return resp, nil
}
