package gohttp

import (
	"errors"
	"net/http"
)

func (c * httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error){
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, errors.New("unable to generate a request")
	}

	allHeaders := c.getRequestHeaders(headers)
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
