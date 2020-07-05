package main

import (
	"clientapi/config"
	"clientapi/entity"
	"clientapi/loader"
	"clientapi/portdomain"
	"clientapi/web"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type store interface {
	Save(ctx context.Context, portDetails entity.PortDetails) error
}

func importFile(fileName string, store store) {
	if fileName == "" {
		log.Printf("[info] no file to load")
		return
	}

	log.Printf("[info] load file: %s", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("[error] file open %s: %v", fileName, err)
		return
	}
	defer file.Close()

	l, err := loader.New(file)
	if err != nil {
		log.Printf("[error] file read: %v", err)
	}

	count := 0
	for row, err := l.NextRow(); err != loader.ErrEOF; row, err = l.NextRow() {
		if err != nil {
			log.Printf("[error] file read: %v", err)
			return
		}

		count++
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		if err := store.Save(ctx, row); err != nil {
			cancel()
			log.Printf("[error] file read: %v", err)
			return
		}
		cancel()
	}
	log.Printf("[info] file load complete, %d rows read and sent", count)
}

func main() {
	conf := config.Load()
	portDomainClient := portdomain.New(conf.PortDomain)

	if err := portDomainClient.Connect(); err != nil {
		log.Fatalf("[fatal] failed to connect to port domain service: %v", err)
	}

	go importFile(conf.FileName, portDomainClient)
	w := web.New(portDomainClient, conf.ServerPort)

	go func() {
		if err := w.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[fatal] server error: %v", err)
		}
	}()

	log.Printf("[info] server started")
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-done
	log.Println("[info] server stopping")
	if err := w.Stop(); err != nil {
		log.Fatalf("[fatal] server shutdown: %v", err)
	}

	if err := portDomainClient.Disconnect(); err != nil {
		log.Fatalf("[fatal] port-domain disconnect: %v", err)
	}
	log.Print("[info] server stopped")
}
