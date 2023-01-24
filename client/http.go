package client

import "mscmsfinder/helper"

func RequestURL(url string) string {
	message := ""

	if !helper.IsValidURL(url) {
		message = "Url is not valid"
		return message
	}

	return message
}
