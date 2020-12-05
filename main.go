package main

import (
	"github.com/korableg/mini-gin/engine"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func shutdown() {
	log.Println("Shutting down Flow...")
	engine.Close()
}

func main() {

	log.Println("Starting Flow...")

	defer shutdown()

	engine.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-quit

}
