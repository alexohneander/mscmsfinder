package client

import (
	"mscmsfinder/helper"
	"mscmsfinder/types"
	"net/http"
)

func TestEndpoint(url string, path string) types.ParseResponse {
	var res types.ParseResponse
	fullUrl := url + path

	if !helper.IsValidURL(fullUrl) {
		res.Message = "Url is not valid"
	} else {
		res = requestEndpoint(fullUrl)
	}

	return res
}

func requestEndpoint(url string) types.ParseResponse {
	var res types.ParseResponse

	httpResp, err := http.Get(url)
	if err != nil {
		res.Message = "Cannot request URL"
		res.Status = 500
	} else {
		res.Message = httpResp.Status
		res.Status = httpResp.StatusCode
		res.Payload.Header = httpResp.Header
		res.Payload.Body = httpResp.Body
	}

	return res
}
