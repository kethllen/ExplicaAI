package configuration

import (
	"context"

	"github.com/kethllen/explicaAi/internal/infrastructure/api"
	"github.com/kethllen/explicaAi/internal/infrastructure/log"
	"github.com/labstack/echo/v4"
)

type Application struct {
	server *echo.Echo
}

func NewApplication() *Application {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	return &Application{
		server: server,
	}
}

func (a *Application) Start() {
	a.registerControllers()
	ctx := context.Background()

	log.LogInfo(ctx, "explicAI is starting on 0.0.0.0:8080")
	log.LogError(ctx, "server fatal error", a.server.Start("0.0.0.0:8080"))

}
func (a *Application) registerControllers() {
	api.NewExplicaServer().Register(a.server)
}
