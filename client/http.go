package client

import (
	"bytes"
	"fmt"
	"mscmsfinder/helper"
	"mscmsfinder/model"
	"mscmsfinder/types"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func TestEndpoint(url string, cms model.CMS) types.ParseResponse {
	var res types.ParseResponse
	fullUrl := url + cms.UrlPath

	if !helper.IsValidURL(fullUrl) {
		res.Message = "Url is not valid"
	} else {
		res = requestEndpoint(fullUrl)
	}

	if res.Payload.Body != "" {
		res = parsePayload(res, cms)
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
		// Close Connection
		defer httpResp.Body.Close()

		// Read Buffer
		buf := new(bytes.Buffer)
		buf.ReadFrom(httpResp.Body)
		content := buf.String()

		// Create Response
		res.Message = httpResp.Status
		res.Status = httpResp.StatusCode
		res.Payload.Header = httpResp.Header
		res.Payload.Body = content
	}

	return res
}

func parsePayload(response types.ParseResponse, cms model.CMS) types.ParseResponse {

	test := html.UnescapeString(response.Payload.Body)
	fmt.Println(test)

	foundCMS := strings.Contains(response.Payload.Body, cms.QueryString)
	if foundCMS {
		response.Message = "Found CMS"
		response.CMS = cms.Title
		response.Payload.Body = ""
		response.Found = true
	}

	return response
}
