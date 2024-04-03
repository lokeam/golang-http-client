package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lokeam/golang-http-client/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	client.DisableTimeouts(true)

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Some-auth-value-987")

	client.SetHeaders(commonHeaders)
	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
}

type User struct {
	FirstName string
	LastName string
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)

	if (err != nil) {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func getUrls() {
	headers := make(http.Header)

	response, err := githubHttpClient.Get("https://api.github.com", headers)

	if (err != nil) {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
