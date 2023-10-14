package server

import (
	"github.com/ryanadiputraa/spotwave/api/internal/oauth/controller"
	"github.com/ryanadiputraa/spotwave/api/internal/oauth/service"
	"github.com/ryanadiputraa/spotwave/api/pkg/spotify"
)

func (s *Server) mapHandlers() {
	oauth := s.fiber.Group("/oauth")

	spotifyUtil := spotify.NewSpotifyUtil()
	service := service.NewOauthService(s.config, spotifyUtil)
	controller.NewOauthController(oauth, s.config, service, spotifyUtil)
}
