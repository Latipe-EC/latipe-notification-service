package scheduleGrpc

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"latipe-notification-service/config"
)

type scheduleServiceGRPCClientImpl struct {
	cfg *config.AppConfig
	cc  grpc.ClientConnInterface
}

func NewDeliveryServiceGRPCClientImpl(config *config.AppConfig) ScheduleServiceClient {
	// Set up a connection to the server.
	log.Info("[GRPC Client] open connection to delivery service")
	conn, err := grpc.Dial(config.GrpcInfrastructure.ScheduleGRPC.Connection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("did not connect: %v", err)
	}

	return &scheduleServiceGRPCClientImpl{
		cfg: config,
		cc:  conn,
	}
}

func (c *scheduleServiceGRPCClientImpl) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*CreateScheduleResponse, error) {
	md := metadata.New(
		map[string]string{"x-api-key": c.cfg.GrpcInfrastructure.ScheduleGRPC.APIGrpcKey},
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	out := new(CreateScheduleResponse)
	err := c.cc.Invoke(ctx, "/ScheduleService/CreateSchedule", in, out, opts...)
	if err != nil {
		log.Errorf("Error invoking CreateSchedule: %v", err)
		return nil, err
	}

	log.Info("[GRPC Client] CreateSchedule is success")
	return out, nil
}
