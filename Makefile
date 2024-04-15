#setup > wire > clean > build > run


SERVICE_NAME = notification_service
WORKER_MAIN_FILE = server_app
BUILD_DIR = $(PWD)/build
GO= go

setup:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest

wire:
	cd internal/ && wire
wire-linux:
	cd internal/ && ~/go/bin/wire

protoc:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./internal/grpc-service/notificationGrpc/notification_service.proto
#linux
# clean build file
cleanl:
	echo "remove bin exe"
	rm -rf $(BUILD_DIR)

# build binary
buildl:
	echo "build binary execute file"
	make wire-linux
	cd cmd/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(WORKER_MAIN_FILE)_linux .

runl:
	make buildl

	echo "Run service application"
	cd $(BUILD_DIR) && ./$(WORKER_MAIN_FILE)_linux


#windows
cleanw:
	echo "remove bin exe"
	rd /s build

buildw:
	echo "build binary execute file"
	make wire
	cd cmd/ &&  $(GO) build -o ..$(BUILD_DIR)/$(WORKER_MAIN_FILE)_win.exe .

runw:
	.\$(BUILD_DIR)\$(WORKER_MAIN_FILE)_win.exe

startw:
	make buildw
	.\$(BUILD_DIR)\$(WORKER_MAIN_FILE)_win.exe