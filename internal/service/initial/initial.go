package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/template-golang/internal/service"
	"github.com/wisaitas/template-golang/pkg/httpx"
)

func init() {
	if err := env.Parse(&service.Config); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}
}

type App struct {
	FiberApp *fiber.App
}

func New() *App {

	app := fiber.New()

	app.Use(httpx.NewLogger(service.Config.Server.Name))

	_ = newMiddleware(app)

	return &App{
		FiberApp: app,
	}
}

func (i *App) Run() {
	go func() {
		i.FiberApp.Listen(fmt.Sprintf(":%d", service.Config.Server.Port))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func (i *App) Close() {
	if err := i.FiberApp.Shutdown(); err != nil {
		log.Fatalf("failed to shutdown fiber app: %v", err)
	}

	log.Println("fiber app shutdown")
}
