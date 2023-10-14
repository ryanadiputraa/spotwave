package service

import (
	"context"

	"github.com/ryanadiputraa/spotwave/api/config"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
	"github.com/ryanadiputraa/spotwave/api/internal/spotify"
	spotifyUtil "github.com/ryanadiputraa/spotwave/api/pkg/spotify"
)

type service struct {
	config      *config.Config
	spotifyUtil spotifyUtil.SpotifyUtil
}

func NewService(config *config.Config, spotifyUtil spotifyUtil.SpotifyUtil) spotify.Usecase {
	return &service{
		config:      config,
		spotifyUtil: spotifyUtil,
	}
}

func (s *service) GetUserInfo(ctx context.Context, accessToken string) (domain.SpotifyUser, error) {
	return s.spotifyUtil.GetUserInfo(accessToken)
}
