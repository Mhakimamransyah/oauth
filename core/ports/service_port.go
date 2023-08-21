package ports

import "github.com/Mhakimamransyah/oauth/core/entities"

type Service interface {

	// register new user
	Register(account int, code string) (*entities.User, error)

	// get detail user by email
	GetDetailUserByEmail(email string) (*entities.User, error)
}
