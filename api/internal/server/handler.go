package server

import (
	oauthController "github.com/ryanadiputraa/spotwave/api/internal/oauth/controller"
	oauthService "github.com/ryanadiputraa/spotwave/api/internal/oauth/service"
	spotifyController "github.com/ryanadiputraa/spotwave/api/internal/spotify/controller"
	spotifyService "github.com/ryanadiputraa/spotwave/api/internal/spotify/service"
	rapidapi "github.com/ryanadiputraa/spotwave/api/pkg/rapid-api"
	spotifyUtil "github.com/ryanadiputraa/spotwave/api/pkg/spotify"
)

func (s *Server) mapHandlers() {
	oauth := s.fiber.Group("/oauth")
	spotify := s.fiber.Group("/api/spotify")

	spotifyUtil := spotifyUtil.NewSpotifyUtil()
	rapidAPI := rapidapi.NewRapidAPI(s.config.APIKey)

	spotifyService := spotifyService.NewService(s.config, spotifyUtil, rapidAPI)
	spotifyController.NewController(spotify, s.config, spotifyService)

	oauthService := oauthService.NewService(s.config, spotifyUtil)
	oauthController.NewController(oauth, s.config, oauthService, spotifyUtil)
}
