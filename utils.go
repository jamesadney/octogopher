package octogopher

import (
	"net/http"
	"net/url"
	"strings"
)

// TODO: Should this be in the Client.Do() method?
// TODO: Should http errors be returned as errors?
// TODO: clean up status code catching stuff
func catchHttpErrors(resp *http.Response) error {
	if resp.StatusCode%400 < 100 {
		return ClientError(resp)
	} else if resp.StatusCode%500 < 100 {
		return ServerError(resp)
	}
	return nil
}

// net/url doesn't replace `/` and `;` because it works on the entire path, not
// just one element.
func encodeUrlPathDelimiters(pathElement string) string {
	r := strings.NewReplacer("/", "%2F", ";", "%3B")
	return r.Replace(pathElement)
}

func urlAppend(base string, elements []string) (string, error) {
	encoded := make([]string, len(elements))
	for i, e := range elements {
		encoded[i] = encodeUrlPathDelimiters(e)
	}
	//TODO: remove trailing slash from base url? or does net/url do that?
	all := make([]string, 1, len(encoded)+1)
	all[0] = base
	all = append(all, encoded...)

	rawURL := strings.Join(all, "/")
	encodedURL, err := url.Parse(rawURL)

	return encodedURL.String(), err
}
