package helper

import (
	"encoding/json"
	"time"

	"github.com/Mhakimamransyah/oauth/config"
	"github.com/Mhakimamransyah/oauth/core/entities"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
	jwt.Claims
}

func (obj *JWTClaims) ConvertToUser() *entities.User {
	return &entities.User{
		Name:  obj.Name,
		Email: obj.Email,
	}
}

func (obj *JWTClaims) BindUser(user *entities.User) *JWTClaims {

	obj.Name = user.Name
	obj.Email = user.Email
	obj.Exp = time.Now().Add(time.Minute * 5).Unix()

	return obj
}

func (obj *JWTClaims) GetToken() (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, obj)
	token, err := jwtToken.SignedString([]byte(config.Config.JwtKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeToken(token string) *entities.User {

	var customClaims JWTClaims

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JwtKey), nil
	})

	if err != nil {
		panic(err)
	}

	jsonString, err := json.Marshal(claims)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(jsonString, &customClaims); err != nil {
		panic(err)
	}

	return customClaims.ConvertToUser()
}
