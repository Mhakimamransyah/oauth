package routes

import (
	"github.com/Mhakimamransyah/oauth/api"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	handler *api.Handler
}

func (router *Router) RegisterRouter(app *fiber.App) {

	// post oauth code
	app.Post("api/oauth", router.handler.Register)

	// get user detail
	app.Get("api/users", api.NewAuthMiddleware(), router.handler.GetDetailUser)

}

func NewRouter(handler *api.Handler) *Router {
	return &Router{
		handler: handler,
	}
}
