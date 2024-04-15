package grpc_service

import (
	"github.com/google/wire"
	"latipe-notification-service/internal/grpc-service/interceptor"
	"latipe-notification-service/internal/grpc-service/notificationGrpc"
)

var Set = wire.NewSet(
	notificationGrpc.NewNotificationGrpcServer,
	interceptor.NewGrpcInterceptor,
)
