package domain

import "fmt"

const (
	SpotifyAccountAPIURL = "https://accounts.spotify.com"
	SpotifyBaseAPIURL    = "https://api.spotify.com/v1"
)

type SpotifyUser struct {
	ID          string            `json:"id"`
	DisplayName string            `json:"display_name"`
	Images      []SpotifyImageURL `json:"images"`
	ProfileURL  string            `json:"href"`
	Email       string            `json:"email"`
	Country     string            `json:"country"`
}

type SpotifyImageURL struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SpotifyPlaylists struct {
	Items    []SpotifyPlaylistItem `json:"items"`
	Limit    int                   `json:"limit"`
	Next     int                   `json:"next"`
	Offset   int                   `json:"offset"`
	Previous int                   `json:"previous"`
	Total    int                   `json:"total"`
}

type SpotifyPlaylistItem struct {
	ID     string            `json:"id"`
	Images []SpotifyImageURL `json:"images"`
	Name   string            `json:"name"`
}

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

type SpotifyOauthError struct {
	ErrorCode        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e *SpotifyOauthError) Error() string {
	return fmt.Sprintf("Error %v : %v", e.ErrorCode, e.ErrorDescription)
}

type SpotifyError struct {
	ErrorDetail SpotifyErrorDetail `json:"error"`
}

type SpotifyErrorDetail struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *SpotifyError) Error() string {
	return fmt.Sprintf("Error %v", e.ErrorDetail.Message)
}
