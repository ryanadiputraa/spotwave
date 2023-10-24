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

func (s *service) GetUserPlaylists(ctx context.Context, accessToken string) (domain.SpotifyPlaylists, error) {
	return s.spotifyUtil.GetUserPlaylist(accessToken)
}

func (s *service) GetPlaylistTracks(ctx context.Context, accessToken, playlistID string) (domain.SpotifyPlaylistTracks, error) {
	return s.spotifyUtil.GetPlaylistTracks(accessToken, playlistID)
}
