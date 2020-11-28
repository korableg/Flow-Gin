package main

import (
	"github.com/korableg/mini-gin/Engine"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func shutdown() {
	log.Println("Shutting down Flow...")
	Engine.Close()
}

func main() {

	log.Println("Starting Flow...")

	defer shutdown()

	Engine.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

}
