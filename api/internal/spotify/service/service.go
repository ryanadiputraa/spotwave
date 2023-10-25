package service

import (
	"context"
	"fmt"

	"github.com/ryanadiputraa/spotwave/api/config"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
	"github.com/ryanadiputraa/spotwave/api/internal/spotify"
	rapidapi "github.com/ryanadiputraa/spotwave/api/pkg/rapid-api"
	spotifyUtil "github.com/ryanadiputraa/spotwave/api/pkg/spotify"
	"github.com/sagikazarmark/slog-shim"
)

type service struct {
	config      *config.Config
	spotifyUtil spotifyUtil.SpotifyUtil
	rapidAPI    rapidapi.RapidAPI
}

func NewService(config *config.Config, spotifyUtil spotifyUtil.SpotifyUtil, rapidAPI rapidapi.RapidAPI) spotify.Usecase {
	return &service{
		config:      config,
		spotifyUtil: spotifyUtil,
		rapidAPI:    rapidAPI,
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

func (s *service) DownloadTrack(ctx context.Context, artists, trackID string) (link string, err error) {
	// TODO: fetch youtube api for track, and use the video id
	data, err := s.rapidAPI.DownloadYoutubeMP3("fshkiZzyF14")
	if err != nil {
		slog.Error(fmt.Sprintf("fail to download: %v", err))
		return "", err
	}
	slog.Info(fmt.Sprintf("download: %v -  %v - %v", data.Title, data.Duration, data.Link))

	return data.Link, nil
}
