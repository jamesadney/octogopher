package octogopher

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
)

func (c *Client) Repos(since int) (repos []*Repo, err error) {
	params := url.Values{}
	params.Set("since", strconv.Itoa(since))

	resp, err := c.Get(REPOS_URL, params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &repos)
	return
}
