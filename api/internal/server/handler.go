package server

import (
	"github.com/ryanadiputraa/spotwave/api/internal/oauth/controller"
)

func (s *Server) mapHandlers() {
	oauth := s.fiber.Group("/oauth")

	controller.NewOauthController(oauth, s.config)
}
