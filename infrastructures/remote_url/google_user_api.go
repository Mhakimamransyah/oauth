package remoteurl

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mhakimamransyah/oauth/config"
	"github.com/Mhakimamransyah/oauth/core/entities"
)

type GoogleUserApi struct {
}

func (obj *GoogleUserApi) FetchUserInformation(token string, user *entities.User) error {

	var client = &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET",
		config.Config.GoogleOauth.UserInformationUrl,
		nil,
	)

	if err != nil {
		return err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	var googleUser = entities.GoogleUserInformation{}
	if err := json.NewDecoder(response.Body).Decode(&googleUser); err != nil {
		return err
	}

	*user = *googleUser.ConvertToUser()
	user.Token = token
	user.Account = entities.GoogleAccount
	user.RegisteredAt = time.Now()

	return nil
}

func NewGoogleUserInformation() *GoogleUserApi {
	return &GoogleUserApi{}
}
