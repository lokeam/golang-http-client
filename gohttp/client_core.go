package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxIdleConnections 	= 5
	defaultResponseTimeout 			= 5 * time.Second
	defaultConnectionTimeout 		= 1 * time.Second
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

	client := c.getHttpClient()

	return client.Do(request)
}

func(c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	c.client = &http.Client{
		Transport: &http.Transport{
			// Value should be configured based on estimated pattern of requests/min
			MaxIdleConnsPerHost: 		c.getMaxIdleConnections(),
			// Max amount of time to wait for a response after sending request
			ResponseHeaderTimeout: 	c.getResponseTimeout(),
			DialContext: 						(&net.Dialer{
				// Max amount of time to wait for any given connection
				Timeout: c.getConnectionTimeout(),
			}).DialContext,
		},
	}
	return c.client
}

func(c *httpClient) getMaxIdleConnections() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func(c *httpClient) getResponseTimeout() time.Duration {
	if c.responseTimeout > 0 {
		return c.responseTimeout
	}

	return defaultResponseTimeout
}

func(c *httpClient) getConnectionTimeout() time.Duration {
	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}

	return defaultConnectionTimeout
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
