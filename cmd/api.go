package main

import (
	"log"

	"github.com/ryanadiputraa/spotwave/api/config"
	"github.com/ryanadiputraa/spotwave/api/internal/server"
)

func main() {
	config, err := config.LoadConfig("config")
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(config)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
