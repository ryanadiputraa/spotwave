package oauth

import (
	"context"

	"github.com/ryanadiputraa/spotwave/api/internal/domain"
)

type Usecase interface {
	Callback(ctx context.Context, code string) (domain.SpotifyAccessTokens, error)
	RefreshToken(ctx context.Context, refreshToken string) (domain.SpotifyRefreshTokens, error)
}
