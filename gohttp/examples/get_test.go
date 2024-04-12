package examples

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/lokeam/golang-http-client/gohttp"
)

func TestGetEndpoints(t *testing.T) {
	gohttp.StartMockServer()

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T)  {
		gohttp.AddMock(gohttp.Mock{
			Method: 	http.MethodGet,
			Url: 			"https://api.github.com",
			Error: 		errors.New("timeout getting github endpoints"),
		})

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("error expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestErrorUnmarshallResponseBody", func(t *testing.T)  {
		gohttp.AddMock(gohttp.Mock{
			Method: 							http.MethodGet,
			Url: 									"https://api.github.com",
			RequestBody: 					`{"current_user_url": 123}`,
			ResponseStatusCode: 	http.StatusOK,
		})

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("error expected")
		}

		if !strings.Contains(err.Error(), "json unmarshal error") {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestNoError", func(t *testing.T)  {
		gohttp.AddMock(gohttp.Mock{
			Method: 							http.MethodGet,
			Url: 									"https://api.github.com",
			RequestBody: 					`{"current_user_url": "https://api.github.com/user"}`,
			ResponseStatusCode: 	http.StatusOK,
		})

		endpoints, err := GetEndpoints()

		if err != nil {
			t.Errorf("Unexpected error: '%s'", err.Error())
		}

		if endpoints == nil {
			t.Error("nil returned when endpoints expected")
		}

		if endpoints.CurrentUser != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})

	endpoints, err := GetEndpoints()

	fmt.Println(err)
	fmt.Println(endpoints)

}