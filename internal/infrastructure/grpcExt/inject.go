package grpcExt

import (
	"github.com/google/wire"
	"latipe-notification-service/internal/infrastructure/grpcExt/scheduleGrpc"
)

var Set = wire.NewSet(
	scheduleGrpc.NewDeliveryServiceGRPCClientImpl,
)
