package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	server "latipe-notification-service/internal"
	"runtime"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Print("Starting server initialization")
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

	endTime := time.Now()
	fmt.Printf("Server initialization completed in %v", endTime.Sub(startTime))

	wg.Wait()
}
