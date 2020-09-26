package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/artzor/tech-test/port-domain/config"
	"github.com/artzor/tech-test/port-domain/service"
	"github.com/artzor/tech-test/port-domain/store"
)

func startServer() {
	conf := config.Load()

	db, err := store.Connect(conf.DBInstance, conf.DBName)
	if err != nil {
		log.Fatalf("[fatal] db connect: %v", err)
	}

	svc := service.New(
		store.New(db),
		conf.ServerPort,
	)

	go func() {
		if err := svc.Start(); err != nil {
			log.Fatalf("[fatal] server start: %v", err)
		}
	}()

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-done
	log.Println("[info] server stopping")
	svc.Stop()
	log.Println("[info] server stopped")
}

func main() {
	startServer()
}
