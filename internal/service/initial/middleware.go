package initial

import (
	"github.com/gofiber/fiber/v2"
	middlewareConfig "github.com/wisaitas/template-golang/internal/service/middleware/config"
)

type middleware struct {
}

func newMiddleware(app *fiber.App) *middleware {
	app.Use(middlewareConfig.Healthz())
	return &middleware{}
}
