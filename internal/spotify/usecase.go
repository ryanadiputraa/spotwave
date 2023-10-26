package spotify

import (
	"context"

	"github.com/ryanadiputraa/spotwave/api/internal/domain"
)

type Usecase interface {
	GetUserInfo(ctx context.Context, accessToken string) (domain.SpotifyUser, error)
	GetUserPlaylists(ctx context.Context, accessToken string) (domain.SpotifyPlaylists, error)
	GetPlaylistTracks(ctx context.Context, accessToken, playlistID string) (domain.SpotifyPlaylistTracks, error)
	DownloadTrack(ctx context.Context, artists, title string) (link string, err error)
}
