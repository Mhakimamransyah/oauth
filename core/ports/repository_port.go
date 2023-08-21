package ports

import (
	"github.com/Mhakimamransyah/oauth/core/entities"
)

type Repository interface {

	// get github access token
	GetAccessToken(account int, code string) (string, error)

	// get github user information
	GetUserInformation(account int, accessToken string) (*entities.User, error)

	// Get user
	FindUserByEmail(email string) (*entities.User, error)

	// store user information to local database
	StoreDetailUser(user *entities.User) error

	// store user information to local database
	IncrUserLoginCount(user *entities.User) error
}
