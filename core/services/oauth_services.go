package services

import (
	"errors"

	"github.com/Mhakimamransyah/oauth/core/entities"
	"github.com/Mhakimamransyah/oauth/core/ports"
	"github.com/Mhakimamransyah/oauth/helper"
)

type OauthServices struct {
	repo ports.Repository
}

func (o *OauthServices) Register(account int, code string) (*entities.User, error) {

	token, err := o.repo.GetAccessToken(account, code)

	if err != nil {
		return nil, err
	}

	user, err := o.repo.GetUserInformation(account, token)

	if err != nil {
		return nil, err
	}

	existingUser, err := o.repo.FindUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	// @todo handling this
	if existingUser.Id != 0 && existingUser.Account != user.Account {
		return nil, errors.New("You're Email Just Registered Using Another Account")
	}

	if existingUser.Id == 0 {
		// insert newly registered user
		return o.createUserResources(true, user)
	}

	return o.createUserResources(false, existingUser)
}

func (o *OauthServices) GetDetailUserByEmail(email string) (*entities.User, error) {
	return o.repo.FindUserByEmail(email)
}

func (o *OauthServices) createUserResources(createNewUser bool, user *entities.User) (*entities.User, error) {

	var err error

	if createNewUser == true {
		if err = o.repo.StoreDetailUser(user); err != nil {
			return nil, err
		}
	} else {
		if err = o.repo.IncrUserLoginCount(user); err != nil {
			return nil, err
		}
	}

	var claims helper.JWTClaims
	user.Token, err = claims.BindUser(user).GetToken()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewOauthServices(repo ports.Repository) *OauthServices {
	return &OauthServices{
		repo: repo,
	}
}
