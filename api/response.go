package api

import (
	"time"

	"github.com/Mhakimamransyah/oauth/core/entities"
)

type SuccessRegisterResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Token string `json:"token"`
}

// Response Success Register new user
func NewSuccessRegisterResponse(user *entities.User) *SuccessRegisterResponse {
	return &SuccessRegisterResponse{
		Email: user.Email,
		Name:  user.Name,
		Image: user.Image,
		Token: user.Token,
	}
}

type SuccessGetDetailUser struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	Account      int       `json:"account"`
	CountLogin   int       `json:"count_login"`
	RegisteredAt time.Time `json:"registered_at"`
}

// Response success get detail user
func NewSuccessGetDetailUser(user *entities.User) *SuccessGetDetailUser {
	return &SuccessGetDetailUser{
		Id:           user.Id,
		Email:        user.Email,
		Name:         user.Name,
		Image:        user.Image,
		Account:      user.Account,
		CountLogin:   user.CountLogin,
		RegisteredAt: user.RegisteredAt,
	}
}
