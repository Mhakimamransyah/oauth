package remoteurl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mhakimamransyah/oauth/config"
	"github.com/Mhakimamransyah/oauth/core/entities"
)

type AccessToken struct {
}

func (a *AccessToken) FetchAccessToken(account int, code string) (string, error) {

	var url string
	var client = &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if account == entities.GithubAccount {
		url = a.getGithubAccessTokenURL(code)
	} else {
		url = a.getGoogleAccessTokenURL(code)
	}

	request, err := http.NewRequestWithContext(ctx, "POST", url, nil)

	if err != nil {
		return "", err
	}

	request.Header.Set("accept", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var accessToken = struct {
		AccessToken string `json:"access_token"`
	}{}

	if err := json.NewDecoder(response.Body).Decode(&accessToken); err != nil {
		return "", err
	}

	return accessToken.AccessToken, nil
}

func (a *AccessToken) getGithubAccessTokenURL(code string) string {

	return fmt.Sprintf(
		config.Config.GithubOauth.AccesTokenUrl+"?client_id=%s&client_secret=%s&code=%s",
		config.Config.GithubOauth.ClientId,
		config.Config.GithubOauth.ClientSecret,
		code,
	)
}

func (a *AccessToken) getGoogleAccessTokenURL(code string) string {

	return fmt.Sprintf(
		config.Config.GoogleOauth.AccesTokenUrl+"?client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=%s&code=%s",
		config.Config.GoogleOauth.ClientId,
		config.Config.GoogleOauth.ClientSecret,
		config.Config.GoogleOauth.RedirectUri,
		"authorization_code",
		code,
	)
}

func NewAccessToken() *AccessToken {
	return &AccessToken{}
}
