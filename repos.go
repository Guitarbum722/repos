package repos

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const baseURL = "https://api.github.com"

// Clienter ..
type Clienter interface {
	RepoCount(string) (int, error)
}

// GithubClient is a http.Client that will provide options specific to Github API requests.
type GithubClient struct {
	hc *http.Client
}

// UserResponse represents the response data from the github user api.
type UserResponse struct {
	PublicRepos int `json:"public_repos"`
}

// RepoCount makes a GET request to the Github '/users' API and
// returns the number of public repositories that the specified User owns.
// If an error occurs, the error will be returned with a repo count of -1
func RepoCount(username string) (int, error) {
	client := &GithubClient{}
	uri := "/users/" + username

	var user UserResponse
	if err := client.call(http.MethodGet, uri, nil, &user); err != nil {
		return -1, err
	}

	return user.PublicRepos, nil
}

func (gc *GithubClient) call(method, uri string, data *bytes.Buffer, result interface{}) error {
	var req *http.Request
	var err error

	endpoint := baseURL + uri

	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(method, endpoint, nil)
		if err != nil {
			return err
		}
	default:
		return errors.New("err - unsupported http method : " + method)
	}

	defer func() { req.Close = true }()

	res, err := gc.hc.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(result)
}
