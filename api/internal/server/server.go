package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/config"
)

type Server struct {
	config *config.Config
	fiber  *fiber.App
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		fiber:  fiber.New(),
	}
}

func (s *Server) Run() error {
	return s.fiber.Listen(s.config.Port)
}
