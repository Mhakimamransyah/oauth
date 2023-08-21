package tables

import (
	"time"

	"github.com/Mhakimamransyah/oauth/core/entities"
)

type User struct {
	Id           int       `gorm:"column:id;primaryKey"`
	Name         string    `gorm:"column:name"`
	Email        string    `gorm:"email"`
	Image        string    `gorm:"image"`
	Account      int       `gorm:"account"`
	CountLogin   int       `gorm:"count_login"`
	RegisteredAt time.Time `gorm:"registered_at"`
	CreatedAt    time.Time `gorm:"created_at"`
	UpdatedAt    time.Time `grom:"updated_at"`
}

func (obj *User) ConvertToUserEntity() *entities.User {
	return &entities.User{
		Id:           obj.Id,
		Name:         obj.Name,
		Email:        obj.Email,
		Image:        obj.Image,
		Account:      obj.Account,
		CountLogin:   obj.CountLogin,
		RegisteredAt: obj.RegisteredAt,
	}
}

func (obj *User) BindNewUser(user *entities.User) {
	obj.Name = user.Name
	obj.Email = user.Email
	obj.Image = user.Image
	obj.Account = user.Account
	obj.CountLogin = user.CountLogin
	obj.RegisteredAt = user.RegisteredAt
	obj.CreatedAt = time.Now()
}
