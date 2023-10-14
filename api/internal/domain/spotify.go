package domain

import "fmt"

const (
	SpotifyAccountAPIURL = "https://accounts.spotify.com"
	SpotifyBaseAPIURL    = "https://api.spotify.com/v1"
)

type SpotifyAccessTokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type SpotifyRefreshTokens struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
}

type SpotifyError struct {
	ErrorCode        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e *SpotifyError) Error() string {
	return fmt.Sprintf("Error %v : %v", e.ErrorCode, e.ErrorDescription)
}
