## Latipe Notification Service

### Description

This is a notification service that push notifications via `Firebase Cloud Message`.
It is built using Go (Fiber and gRPC).

### Technologies

- Go 1.20
- Fiber (v2)
- gRPC
- Firebase Cloud Message
- MongoDB
- RabbitMQ
- Redis
- Docker

### Features

- Send notification to many devices of a user
- Send campaign notification to many users
- Get notification list of a user
- Get notification detail
- Mark notification as read
- Mark all notifications as read
- Delete all notifications
- [Admin] Health check and metrics

### How to run

1. Clone this repository.
2. Assure that you have run MongoDB, RabbitMQ, and Redis and you have the firebase config file.
3. Update your config in `config/config.yml` and put the firebase config file in same directory.
4. Having two way to run this service:
    - Run using `Dockerfile` to build the image and run the container (expose port 5020 and 6020)
    - Run using `Makefile` (recommended)
        - `make buildw` to build the service for windows
        - `make buildl` to build the service for linux
        - And use `make runw` or `make runl` to run the service

### Documentation

- RestAPI port: 5020
- gRPC port: 6020
- API URL: `/api/v1/notifications/*`
- Proto file is located in `/internal/grpc-service/notificationGrpc/notification_service.proto`
    - Put `x-api-key: 123123` in metadata to access the gRPC service
- Metrics: `/metrics`
- Health check: `/health`
  - Readiness check: `/readiness`
  - Liveness check: `/liveness`
- Dashboard: `/fiber/dashboard`

<hr>

### Development by tdat.it2k2