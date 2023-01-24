package client

import (
	"mscmsfinder/helper"
	"net/http"
)

func TestEndpoint(url string, path string) string {
	message := ""
	fullUrl := url + path

	if !helper.IsValidURL(fullUrl) {
		message = "Url is not valid"
	}

	resp, err := http.Get(fullUrl)
	if err != nil {
		message = "Cannot request URL"
	}

	if resp.StatusCode == 200 {
		message = "Request was successfull"
	}

	return message
}
