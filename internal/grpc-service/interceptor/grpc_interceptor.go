package interceptor

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"latipe-notification-service/config"
	"time"
)

type GrpcInterceptor struct {
	cfg *config.AppConfig
}

func NewGrpcInterceptor(cfg *config.AppConfig) *GrpcInterceptor {
	return &GrpcInterceptor{cfg: cfg}
}

func (mid *GrpcInterceptor) MiddlewareUnaryRequest(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	startTime := time.Now()

	//validate api-key
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if len(md["x-api-key"]) == 0 || md["x-api-key"][0] != mid.cfg.GRPC.APIKey {
			return nil, status.Error(codes.PermissionDenied, "Permission denied")
		}
	}

	data, err := handler(ctx, req)
	if err != nil {
		log.Errorf("[gRPC server] unary request: %v", err)
	}
	endTime := time.Now()
	log.Infof("    [gRPC server] | %v | unary request: %v", endTime.Sub(startTime), info.FullMethod)

	return data, err
}
