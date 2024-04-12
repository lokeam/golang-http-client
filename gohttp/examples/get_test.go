package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/lokeam/golang-http-client/gohttp"
)

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T)  {
		mock := gohttp.Mock{
			Method: 	http.MethodGet,
			Url: 			"https://api.github.com",
			Error: 		errors.New("timeout getting github endpoints"),
		}

		endpoints, err := GetEndpoints()
	})

	t.Run("TestErrorUnmarshallResponseBody", func(t *testing.T)  {
		mock := gohttp.Mock{
			Method: 							http.MethodGet,
			Url: 									"https://api.github.com",
			RequestBody: 					`{"current_user_url": 123}`,
			ResponseStatusCode: 	http.StatusOK,
		}

		endpoints, err := GetEndpoints()
	})

	t.Run("TestNoError", func(t *testing.T)  {
		mock := gohttp.Mock{
			Method: 							http.MethodGet,
			Url: 									"https://api.github.com",
			RequestBody: 					`{"current_user_url": "https://api.github.com/user"}`,
			ResponseStatusCode: 	http.StatusOK,
		}

		endpoints, err := GetEndpoints()
	})

	endpoints, err := GetEndpoints()

	fmt.Println(err)
	fmt.Println(endpoints)

}