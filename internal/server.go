//go:build wireinject
// +build wireinject

package server

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/wire"
	"latipe-notification-service/config"
	"latipe-notification-service/internal/adapter"
	"latipe-notification-service/internal/domain/repositories"
	"latipe-notification-service/internal/handler"
	"latipe-notification-service/internal/middleware"
	"latipe-notification-service/internal/router"
	"latipe-notification-service/internal/router/notifyRouter"
	"latipe-notification-service/internal/service"
	"latipe-notification-service/pkgUtils/db/mongodb"
	"latipe-notification-service/pkgUtils/fcm"
	"latipe-notification-service/pkgUtils/rabbitclient"
)

type Application struct {
	fiberApp  *fiber.App
	appConfig *config.AppConfig
}

func (app Application) FiberApp() *fiber.App {
	return app.fiberApp
}

func (app Application) AppConfig() *config.AppConfig {
	return app.appConfig
}

func New() (*Application, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		config.Set,
		mongodb.Set,
		fcm.Set,
		rabbitclient.Set,
		repositories.Set,
		service.Set,
		adapter.Set,
		handler.Set,
		middleware.Set,
		router.Set,
	)))
}

func NewServer(
	cfg *config.AppConfig,
	notifyRouter notifyRouter.NotificationRouter,
) *Application {

	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		JSONDecoder:  sonic.Unmarshal,
		JSONEncoder:  sonic.Marshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5500",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	prometheus := fiberprometheus.New("notification-service")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Initialize default config
	app.Use(logger.New())

	app.Get("", func(ctx *fiber.Ctx) error {
		s := struct {
			Message string `json:"message"`
			Version string `json:"version"`
		}{
			Message: "notification service was developed by tdatIT",
			Version: "v1.0.0",
		}
		return ctx.JSON(s)
	})
	api := app.Group("/api")
	v1 := api.Group("/v1")
	notifyRouter.Init(&v1)

	return &Application{
		appConfig: cfg,
		fiberApp:  app,
	}
}
