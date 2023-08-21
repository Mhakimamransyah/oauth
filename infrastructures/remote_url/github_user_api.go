package remoteurl

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mhakimamransyah/oauth/config"
	"github.com/Mhakimamransyah/oauth/core/entities"
)

type GithubUserApi struct{}

func (a *GithubUserApi) FetchUserInformation(token string, user *entities.User) error {

	var client = &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET",
		config.Config.GithubOauth.UserInformationUrl,
		nil,
	)

	if err != nil {
		return err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "token "+token)

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	var githubUser = entities.GithubUserInformation{}
	if err := json.NewDecoder(response.Body).Decode(&githubUser); err != nil {
		return err
	}

	*user = *githubUser.ConvertToUser()
	user.Token = token
	user.Account = entities.GithubAccount
	user.RegisteredAt = time.Now()

	return nil
}

func NewGithubUserInformation() *GithubUserApi {
	return &GithubUserApi{}
}
