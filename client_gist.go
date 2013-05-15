package octogopher

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"time"
)

// Default time value is ignored in methods
var NIL_TIME = time.Time{}

func (c *Client) UserGists(username string, since time.Time) (gists []*Gist, err error) {
	params := url.Values{}

	if since != NIL_TIME {
		// FIXME: Github uses ISO 8601. Are there going to be differences?
		params.Set("since", since.Format(time.RFC3339))
	}

	urlStr, err := urlAppend(USERS_URL, []string{username, "gists"})
	if err != nil {
		return nil, err
	}

	resp, err := c.Get(urlStr, params)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &gists)

	return
}
