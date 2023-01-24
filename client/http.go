package client

import (
	"mscmsfinder/helper"
	"net/http"
)

func RequestURL(url string) string {
	message := ""

	if !helper.IsValidURL(url) {
		message = "Url is not valid"
	}

	resp, err := http.Get(url)
	if err != nil {
		message = "Cannot request URL"
	}

	if resp.StatusCode == 200 {
		message = "Request was successfull"
	}

	return message
}
