//go:build wireinject
// +build wireinject

package server

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/google/wire"
	"github.com/hellofresh/health-go/v5"
	"google.golang.org/grpc"
	"latipe-notification-service/config"
	"latipe-notification-service/internal/adapter"
	"latipe-notification-service/internal/domain/repositories"
	grpc_service "latipe-notification-service/internal/grpc-service"
	"latipe-notification-service/internal/grpc-service/interceptor"
	"latipe-notification-service/internal/grpc-service/notificationGrpc"
	"latipe-notification-service/internal/handler"
	"latipe-notification-service/internal/middleware"
	"latipe-notification-service/internal/router"
	"latipe-notification-service/internal/router/notifyRouter"
	"latipe-notification-service/internal/service"
	"latipe-notification-service/internal/subs"
	"latipe-notification-service/internal/subs/notifySubs"
	"latipe-notification-service/pkgUtils/db/mongodb"
	"latipe-notification-service/pkgUtils/fcm"
	healthService "latipe-notification-service/pkgUtils/healthcheck"
	"latipe-notification-service/pkgUtils/rabbitclient"
	responses "latipe-notification-service/pkgUtils/util/response"
)

type Application struct {
	fiberApp   *fiber.App
	grpcApp    *grpc.Server
	appConfig  *config.AppConfig
	notifySubs *notifySubs.NotifyToUserSubs
}

func (app Application) FiberApp() *fiber.App {
	return app.fiberApp
}

func (app Application) AppConfig() *config.AppConfig {
	return app.appConfig
}

func (app Application) GRPCServer() *grpc.Server {
	return app.grpcApp
}

func (app Application) NotifyToUserSubs() *notifySubs.NotifyToUserSubs {
	return app.notifySubs
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
		subs.Set,
		handler.Set,
		middleware.Set,
		router.Set,
		grpc_service.Set,
	)))
}

func NewServer(
	cfg *config.AppConfig,
	notifyRouter notifyRouter.NotificationRouter,
	//grpc service
	notificationGrpcServ notificationGrpc.NotificationServiceServer,
	unaryInterceptor *interceptor.GrpcInterceptor,
	//subs
	notifySubs *notifySubs.NotifyToUserSubs,
) *Application {

	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		JSONDecoder:  sonic.Unmarshal,
		JSONEncoder:  sonic.Marshal,
		ErrorHandler: responses.CustomErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5500,http://127.0.0.1:5173",
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods: "GET,HEAD,OPTIONS,POST,PUT",
	}))

	//providing basic authentication for metrics endpoints
	basicAuthConfig := basicauth.Config{
		Users: map[string]string{
			"admin": "123123",
		},
	}

	// Fiber prometheus
	prometheus := fiberprometheus.New("promotion-services")
	prometheus.RegisterAt(app, "/metrics", basicauth.New(basicAuthConfig))
	app.Use(prometheus.Middleware)
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		s := struct {
			Message string `json:"message"`
			Version string `json:"version"`
		}{
			Message: "Notification services was developed by TienDat",
			Version: "v1.0.1",
		}
		return ctx.JSON(s)
	})

	// Healthcheck
	h, _ := healthService.NewHealthCheckService(cfg)
	app.Get("/health", basicauth.New(basicAuthConfig), adaptor.HTTPHandlerFunc(h.HandlerFunc))
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/liveness",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			result := h.Measure(c.Context())
			return result.Status == health.StatusOK
		},
		ReadinessEndpoint: "/readiness",
	}))

	// healthcheck
	h, _ = healthService.NewHealthCheckService(cfg)
	app.Get("/health", basicauth.New(basicAuthConfig), adaptor.HTTPHandlerFunc(h.HandlerFunc))
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/liveness",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			result := h.Measure(c.Context())
			return result.Status == health.StatusOK
		},
		ReadinessEndpoint: "/readiness",
	}))

	//fiber dashboard
	app.Get("/fiber/dashboard", basicauth.New(basicAuthConfig),
		monitor.New(monitor.Config{Title: "Notification Services Metrics Page (Fiber)"}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	notifyRouter.Init(&v1)

	//init grpc service
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor.MiddlewareUnaryRequest))
	notificationGrpc.RegisterNotificationServiceServer(grpcServer, notificationGrpcServ)

	return &Application{
		appConfig:  cfg,
		fiberApp:   app,
		grpcApp:    grpcServer,
		notifySubs: notifySubs,
	}
}
