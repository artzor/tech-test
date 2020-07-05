package main

import (
	"clientapi/config"
	"log"
)

func main() {
	cfg := config.Load()
	log.Printf("+%v", cfg)
}
