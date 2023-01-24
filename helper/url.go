package helper

import (
	"net/url"
)

func IsValidURL(u string) bool {
	validUrl, err := url.ParseRequestURI(u)

	if err != nil {
		return false
	}

	if validUrl != nil {
		return true
	}

	return false
}
