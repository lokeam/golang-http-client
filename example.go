package main

import (
	"fmt"

	"github.com/lokeam/golang-http-client/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(6).
		Build()

	return client
}

func main() {
	getUrls()
}

type User struct {
	FirstName string
	LastName string
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if (err != nil) {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())
}
