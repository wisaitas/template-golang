package initial

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/template-golang/internal/service"
	middlewareConfig "github.com/wisaitas/template-golang/internal/service/middleware/config"
	"github.com/wisaitas/template-golang/pkg/httpx"
)

type middleware struct {
}

func newMiddleware(app *fiber.App) *middleware {
	app.Use(middlewareConfig.Healthz())
	app.Use(httpx.NewLogger(service.Config.Server.Name))
	return &middleware{}
}
