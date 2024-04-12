package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	s.fiber.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	s.mapHandlers()
	return s.fiber.Listen(fmt.Sprintf(":%v", s.config.Server.Port))
}
