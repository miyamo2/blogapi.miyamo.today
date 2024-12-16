package utils

import (
	"net/url"
)

func MustURLParse(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return *u
}
