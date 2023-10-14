package server

import (
	"github.com/ryanadiputraa/spotwave/api/internal/oauth/controller"
	"github.com/ryanadiputraa/spotwave/api/pkg/spotify"
)

func (s *Server) mapHandlers() {
	oauth := s.fiber.Group("/oauth")

	spotifyUtil := spotify.NewSpotifyUtil()

	controller.NewOauthController(oauth, s.config, spotifyUtil)
}
