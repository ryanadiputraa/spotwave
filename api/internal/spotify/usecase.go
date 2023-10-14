package spotify

import (
	"context"

	"github.com/ryanadiputraa/spotwave/api/internal/domain"
)

type Usecase interface {
	GetUserInfo(ctx context.Context, accessToken string) (domain.SpotifyUser, error)
}
