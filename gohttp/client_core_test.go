package gohttp

import (
	"fmt"
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

	if finalHeaders.Get("X-Request-Id") != "TEST-ABC" {
		t.Error("invalid request id received")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}

	if finalHeaders.Get("User-Agent") != "custom-http-client" {
		t.Error("invalid user agent received")
	}
}


func TestGetRequestBody(t* testing.T) {
	client := httpClient{}
	t.Run("WithoutBodyNilResponse", func(t* testing.T){
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing nil body")
		}

		if body != nil {
			t.Error("no body expected when passing nil body")
		}
	})

	t.Run("BodyWithJson", func(t* testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)

		fmt.Println(err)
		fmt.Println(string(body))

		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("WithBodyWithXml", func(t* testing.T) {

	})
	t.Run("WithBodyWithJsonAsDefault", func(t* testing.T) {

	})



}
