package service

import (
	"net/url"

	"github.com/ryanadiputraa/spotwave/api/config"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
	"github.com/ryanadiputraa/spotwave/api/internal/oauth"
	"github.com/ryanadiputraa/spotwave/api/pkg/spotify"
)

type service struct {
	config      *config.Config
	spotifyUtil spotify.SpotifyUtil
}

func NewOauthService(config *config.Config, spotifyUtil spotify.SpotifyUtil) oauth.Usecase {
	return &service{
		config:      config,
		spotifyUtil: spotifyUtil,
	}
}

func (s *service) Callback(code string) (domain.SpotifyAccessTokens, error) {
	redirect, _ := url.ParseRequestURI(s.config.RedirectURI)
	return s.spotifyUtil.Callback(s.config.ClientID, s.config.ClientSecret, code, redirect.String())
}
