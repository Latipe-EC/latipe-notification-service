package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	server "latipe-notification-service/internal"
	"net"
	"runtime"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println("\n======== Starting server initialization ========")
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPU cores: %d\n", numCPU)
	serv, err := server.New()
	if err != nil {
		log.Fatalf("%s", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := serv.FiberApp().Listen(serv.AppConfig().Server.RestAPIPort); err != nil {
			log.Fatalf("Error: %s", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Infof("Start grpc server on port: localhost%v", serv.AppConfig().Server.GrpcPort)
		lis, err := net.Listen("tcp", serv.AppConfig().Server.GrpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v\n", err)
		}

		if err := serv.GRPCServer().Serve(lis); err != nil {
			log.Infof("%s", err)
		}
	}()

	wg.Add(1)
	go func() {
		if err := serv.NotifyToUserSubs().ListenNotificationMessage(&wg); err != nil {
			log.Fatalf("Error: %s", err)
		}
	}()

	wg.Add(1)
	go func() {
		endTime := time.Now()
		fmt.Printf("\n======== Server initialization completed in %v ========\n", endTime.Sub(startTime))
		wg.Done()
	}()

	wg.Wait()
}
