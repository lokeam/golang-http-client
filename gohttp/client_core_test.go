package gohttp

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

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

	t.Run("WithBodyWithXML", func(t *testing.T) {
    // Define a struct with exported fields and a root element for XML marshaling.
    requestBody := struct {
        XMLName xml.Name `xml:"Root"`
        Item1   string   `xml:"Item1"`
        Item2   string   `xml:"Item2"`
    }{
        Item1: "one",
        Item2: "two",
    }

    body, err := client.getRequestBody("application/xml", requestBody)

    if err != nil {
        t.Errorf("no error expected when marshaling struct as xml, got %v", err)
    }

    // Note: The expected XML string must include the root element and possibly XML declaration.
    // The actual output might include the XML declaration and spaces. Adjust accordingly.
    expectedBody := `<Root><Item1>one</Item1><Item2>two</Item2></Root>`
    // Using strings.Contains to allow flexibility in how the XML is formatted and encoded.
    if !strings.Contains(string(body), expectedBody) {
        t.Errorf("invalid xml body obtained, expected %s, got %s", expectedBody, string(body))
    }
})

	t.Run("WithBodyWithJsonAsDefault", func(t *testing.T) {
    // Testing with no contentType
    requestBody := map[string]string{"key": "value"}

    body, err := client.getRequestBody("", requestBody)

    if err != nil {
        t.Errorf("no error expected when defaulting to JSON marshaling, got %v", err)
    }

    expectedBody := `{"key":"value"}`
    if string(body) != expectedBody {
        t.Errorf("invalid json body obtained, expected %s, got %s", expectedBody, string(body))
    }

    // Testing with unsupported contentType
    body, err = client.getRequestBody("application/unsupported", requestBody)

    if err != nil {
        t.Errorf("no error expected when defaulting to JSON marshaling with unsupported contentType, got %v", err)
    }

    // The expectedBody remains the same as JSON is the default
    if string(body) != expectedBody {
        t.Errorf("invalid json body obtained with unsupported contentType, expected %s, got %s", expectedBody, string(body))
    }
})

}
