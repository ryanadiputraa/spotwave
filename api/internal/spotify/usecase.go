package spotify

import (
	"context"

	"github.com/ryanadiputraa/spotwave/api/internal/domain"
)

type Usecase interface {
	GetUserInfo(ctx context.Context, accessToken string) (domain.SpotifyUser, error)
	GetUserPlaylists(ctx context.Context, accessToken string) (domain.SpotifyPlaylists, error)
}
