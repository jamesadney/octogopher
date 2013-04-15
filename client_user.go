package octogopher

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
)

func (c *Client) Users(since int) (users []*User, err error) {
	params := url.Values{}
	params.Set("since", strconv.Itoa(since))

	resp, err := c.Get(USERS_URL, params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &users)

	return
}

func (c *Client) AuthenticatedUser() (*User, error) {
	return c.getUser(AUTHUSER_URL)
}

func (c *Client) User(username string) (*User, error) {
	urlStr, err := urlAppend(USERS_URL, []string{username})
	if err != nil {
		return nil, err
	}
	return c.getUser(urlStr)
}

func (c *Client) getUser(urlStr string) (*User, error) {
	resp, err := c.Get(urlStr, nil)
	if err != nil {
		return &User{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	user := new(User)
	json.Unmarshal(body, user)

	return user, nil
}
