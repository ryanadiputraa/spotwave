package domain

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
