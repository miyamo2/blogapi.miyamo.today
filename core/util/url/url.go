package url

import "net/url"

func MustParseURL(s string) *url.URL {
	u, _ := url.Parse(s)
	return u
}
