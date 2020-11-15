package main

import (
	"github.com/korableg/mini-gin/Engine"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func shutdown() {
	log.Println("Shutting down mini...")

	Engine.Close()
	// TODO Save messages, nodes, routers to disk
}

func main() {

	log.Println("Starting mini...")

	defer shutdown()

	Engine.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

}
