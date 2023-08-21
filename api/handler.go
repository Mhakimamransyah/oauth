package api

import (
	"net/http"

	"github.com/Mhakimamransyah/oauth/core/entities"
	"github.com/Mhakimamransyah/oauth/core/ports"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services ports.Service
}

func (handler *Handler) GetDetailUser(c *fiber.Ctx) error {

	user := c.Locals("session").(*entities.User)

	existingUser, err := handler.services.GetDetailUserByEmail(user.Email)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(NewSuccessGetDetailUser(existingUser))
}

func (handler *Handler) Register(c *fiber.Ctx) error {

	var oauthSpec = &OauthSpec{}

	if err := c.BodyParser(oauthSpec); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"msg": err.Error(),
		})
	}

	user, err := handler.services.Register(oauthSpec.Account, oauthSpec.Code)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(NewSuccessRegisterResponse(user))
}

func NewHandler(service ports.Service) *Handler {
	return &Handler{
		services: service,
	}
}
