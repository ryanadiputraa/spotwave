package oauth

import "github.com/ryanadiputraa/spotwave/api/internal/domain"

type Usecase interface {
	Callback(code string) (domain.SpotifyAccessTokens, error)
}
