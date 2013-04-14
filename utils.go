package octogopher

import (
	"net/url"
	"strings"
)

// net/url doesn't replace `/` and `;` because it works on the entire path, not
// just one element.
func encodeURLPathDelimiters(pathElement string) string {
	r := strings.NewReplacer("/", "%%2F", ";", "%%3B")
	return r.Replace(pathElement)
}

func urlAppend(base string, elements []string) (string, error) {
	encoded := make([]string, len(elements))
	for i, e := range elements {
		encoded[i] = encodeURLPathDelimiters(e)
	}
	//TODO: remove trailing slash from base url? or does net/url do that?
	all := make([]string, 1, len(encoded)+1)
	all[0] = base
	all = append(all, encoded...)

	rawURL := strings.Join(all, "/")
	encodedURL, err := url.Parse(rawURL)

	return encodedURL.String(), err
}
