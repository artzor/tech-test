package main

import (
	"log"
	"portdomain/config"
)
func main() {
	cfg := config.Load()
	log.Printf("%+v", cfg)
}
