package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Mhakimamransyah/oauth/api"
	"github.com/Mhakimamransyah/oauth/config"
	"github.com/Mhakimamransyah/oauth/core/repositories"
	"github.com/Mhakimamransyah/oauth/core/services"
	"github.com/Mhakimamransyah/oauth/infrastructures/databases/mysql"
	remoteurl "github.com/Mhakimamransyah/oauth/infrastructures/remote_url"
	"github.com/Mhakimamransyah/oauth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.LoadConfig()

	mysql.ConnectMysql()

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	repo := repositories.NewOauthRepository(
		remoteurl.NewAccessToken(),
		remoteurl.NewGithubUserInformation(),
		remoteurl.NewGoogleUserInformation(),
		mysql.MysqlInstance,
	)

	service := services.NewOauthServices(repo)

	routes.NewRouter(
		api.NewHandler(service),
	).RegisterRouter(app)

	app.Static("/css", "public/css")
	app.Static("/js", "public/js")
	app.Static("/web", "public/pages")

	app.Get("/web/*", func(c *fiber.Ctx) error {
		return c.SendFile("public/pages/index.html")
	})

	defer func() {
		mysql.DisconnectMysql()
	}()

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	// wait for interrupt signal
	<-quit

}
