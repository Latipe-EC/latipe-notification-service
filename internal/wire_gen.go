// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
	"github.com/hellofresh/health-go/v5"
	"google.golang.org/grpc"
	"latipe-notification-service/config"
	"latipe-notification-service/internal/adapter/authserv"
	"latipe-notification-service/internal/domain/repositories/notifyRepos"
	"latipe-notification-service/internal/domain/repositories/userDeviceRepos"
	"latipe-notification-service/internal/grpc-service/interceptor"
	"latipe-notification-service/internal/grpc-service/notificationGrpc"
	"latipe-notification-service/internal/handler/notifyHandler"
	"latipe-notification-service/internal/middleware"
	"latipe-notification-service/internal/router/notifyRouter"
	"latipe-notification-service/internal/service/notifyService"
	"latipe-notification-service/pkgUtils/db/mongodb"
	"latipe-notification-service/pkgUtils/fcm"
	"latipe-notification-service/pkgUtils/healthcheck"
	"latipe-notification-service/pkgUtils/util/response"
)

// Injectors from server.go:

func New() (*Application, error) {
	appConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	mongoClient, err := mongodb.OpenMongoDBConnection(appConfig)
	if err != nil {
		return nil, err
	}
	notificationRepository := notifyRepos.NewNotificationRepository(mongoClient)
	userDeviceRepository := userDeviceRepos.NewUserDeviceRepository(mongoClient)
	firebaseCloudMessage := fcm.NewFirebaseSDK(appConfig)
	notificationService := notifyService.NewNotificationService(notificationRepository, userDeviceRepository, firebaseCloudMessage)
	notifyHandlerNotifyHandler := notifyHandler.NewNotifyHandler(notificationService)
	authService := authserv.NewAuthService(appConfig)
	authMiddleware := middleware.NewAuthMiddleware(authService, appConfig)
	notificationRouter := notifyRouter.NewNotificationRouter(notifyHandlerNotifyHandler, authMiddleware)
	notificationServiceServer := notificationGrpc.NewNotificationGrpcServer(notificationService)
	grpcInterceptor := interceptor.NewGrpcInterceptor(appConfig)
	application := NewServer(appConfig, notificationRouter, notificationServiceServer, grpcInterceptor)
	return application, nil
}

// server.go:

type Application struct {
	fiberApp  *fiber.App
	grpcApp   *grpc.Server
	appConfig *config.AppConfig
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

func NewServer(
	cfg *config.AppConfig, notifyRouter2 notifyRouter.NotificationRouter,

	notificationGrpcServ notificationGrpc.NotificationServiceServer,
	unaryInterceptor *interceptor.GrpcInterceptor,
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

	basicAuthConfig := basicauth.Config{
		Users: map[string]string{
			"admin": "123123",
		},
	}

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

	app.Get("/fiber/dashboard", basicauth.New(basicAuthConfig), monitor.New(monitor.Config{Title: "Notification Services Metrics Page (Fiber)"}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	notifyRouter2.
		Init(&v1)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor.MiddlewareUnaryRequest))
	notificationGrpc.RegisterNotificationServiceServer(grpcServer, notificationGrpcServ)

	return &Application{
		appConfig: cfg,
		fiberApp:  app,
		grpcApp:   grpcServer,
	}
}
