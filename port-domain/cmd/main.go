package main

import (
	"log"
	"portdomain/config"
	"portdomain/store"
)

func startServer() {
	conf := config.Load()

	_, err := store.Connect(conf.DBInstance, conf.DBName)
	if err != nil {
		log.Fatalf("[fatal] db connect: %v", err)
	}
}

func main() {
	startServer()
}
