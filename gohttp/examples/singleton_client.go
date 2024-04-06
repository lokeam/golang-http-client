package examples

import (
	"time"

	"github.com/lokeam/golang-http-client/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(5*time.Second).
		SetResponseTimeout(3*time.Second).
		Build()
	return client
}
