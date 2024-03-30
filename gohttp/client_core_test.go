package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	client := httpClient{}

	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "custom-http-client")

	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "TEST-ABC")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	if len(finalHeaders) != 3 {
		t.Error("error: 3 headers expected")
	}
}
