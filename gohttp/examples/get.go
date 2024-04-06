package examples

import "fmt"

type Endpoints struct {
	CurrentUser 			string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl 		string `json:"respository_url"`
}

func GetEndpoints() (*Endpoints, error){
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status Code: %d", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status()))
	fmt.Println(fmt.Sprintf("Body: %s\n", response.String()))

	var endpoints Endpoints
	if err := response.UnmarshalJSON(&endpoints); err != nil {
		return nil, err
	}

	fmt.Printf(fmt.Sprintf("Repositories URL: %s", endpoints.RepositoryUrl))
	return &endpoints, nil
}
