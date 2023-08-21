package api

import (
	"strings"

	"github.com/Mhakimamransyah/oauth/config"
	"github.com/Mhakimamransyah/oauth/helper"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware() fiber.Handler {

	return jwtware.New(

		jwtware.Config{
			SuccessHandler: func(ctx *fiber.Ctx) error {

				authorization := ctx.GetReqHeaders()["Authorization"]
				token := strings.Split(authorization, " ")[1]

				ctx.Locals("session", helper.DecodeToken(token))

				return ctx.Next()
			},
			SigningKey: jwtware.SigningKey{
				Key: []byte(config.Config.JwtKey),
			},
			Filter: func(ctx *fiber.Ctx) bool {
				// middleware activated
				return false
			},
		},
	)

}
