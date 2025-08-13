package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kethllen/explicaAi/configuration"
	"github.com/kethllen/explicaAi/internal/infrastructure/log"
)

func main() {

	go configuration.NewApplication().Start()
	shutDown()
}

func shutDown() {
	signalShutdown := make(chan os.Signal, 2)
	signal.Notify(signalShutdown, syscall.SIGINT, syscall.SIGTERM)
	switch <-signalShutdown {
	case syscall.SIGINT:
		log.LogInfo(context.Background(), "SIGINT signal, explicAI is stopping....")
	case syscall.SIGTERM:
		log.LogInfo(context.Background(), "SIGTERM signal, explicAI is stopping....")
	}
}
