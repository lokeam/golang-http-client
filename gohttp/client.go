package gohttp

import "net/http"

type httpClient struct {}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

func(c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func(c *httpClient) Post() {}

func(c *httpClient) Put() {}

func(c *httpClient) Patch() {}

func(c *httpClient) Delete() {}
