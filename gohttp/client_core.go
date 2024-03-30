package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

func (c* httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}

func (c * httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error){
	client := http.Client{}

	allHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(allHeaders.Get("Content-Type"), body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, errors.New("unable to generate a request")
	}

	request.Header = allHeaders

	return client.Do(request)
}


func(c *httpClient) getRequestHeaders(customHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers
	for headerName, value := range c.Headers {
		if len(value) > 0 {
		  result.Set(headerName, value[0])
		}
	}

	// Add custom headers
	for headerName, value := range customHeaders {
		if len(value) > 0 {
		  result.Set(headerName, value[0])
		}
	}

	return result
}
