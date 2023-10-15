package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
	"github.com/sagikazarmark/slog-shim"
)

type SpotifyUtil interface {
	Login(ctx *fiber.Ctx, clientID, redirectURI, state string) error
	Callback(clientID, clientSecret, code, redirectURI string) (domain.SpotifyAccessTokens, error)
	RefreshToken(clientID, clientSecret, refreshToken string) (domain.SpotifyRefreshTokens, error)
	GetUserInfo(accessToken string) (domain.SpotifyUser, error)
	GetUserPlaylist(accessToken, userID string) (playlists domain.SpotifyPlaylists, err error)
}

type spotify struct{}

func NewSpotifyUtil() SpotifyUtil {
	return &spotify{}
}

func (s *spotify) Login(ctx *fiber.Ctx, clientID, redirectURI, state string) error {
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", clientID)
	params.Add("scope", "user-read-private user-read-email")
	params.Add("redirect_uri", redirectURI)
	params.Add("state", state)
	params.Add("show_dialog", "true")

	u, _ := url.ParseRequestURI(domain.SpotifyAccountAPIURL)
	u.Path = "/authorize"
	u.RawQuery = params.Encode()

	return ctx.Redirect(u.String(), http.StatusTemporaryRedirect)
}

func (s *spotify) Callback(clientID, clientSecret, code, redirectURI string) (tokens domain.SpotifyAccessTokens, err error) {
	u, _ := url.ParseRequestURI(domain.SpotifyAccountAPIURL)
	u.Path = "/api/token"

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		slog.Error("http req: ", err.Error())
		return
	}
	req.Header.Set("Authorization", "basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", clientID, clientSecret))))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("exec http client: ", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var reqErr domain.SpotifyOauthError
		if err = json.NewDecoder(resp.Body).Decode(&reqErr); err != nil {
			slog.Error("decode resp body: ", err.Error())
			return
		}
		return tokens, &reqErr
	}

	err = json.NewDecoder(resp.Body).Decode(&tokens)
	if err != nil {
		slog.Error("decode: ", err.Error())
	}
	return
}

func (s *spotify) RefreshToken(clientID, clientSecret, refreshToken string) (tokens domain.SpotifyRefreshTokens, err error) {
	u, _ := url.ParseRequestURI(domain.SpotifyAccountAPIURL)
	u.Path = "/api/token"

	data := url.Values{}
	data.Set("refresh_token", refreshToken)
	data.Set("grant_type", "refresh_token")

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		slog.Error("http req: ", err.Error())
		return
	}
	req.Header.Set("Authorization", "basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", clientID, clientSecret))))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("exec http client: ", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var reqErr domain.SpotifyOauthError
		if err = json.NewDecoder(resp.Body).Decode(&reqErr); err != nil {
			slog.Error("decode resp body: ", err.Error())
			return
		}
		return tokens, &reqErr
	}

	err = json.NewDecoder(resp.Body).Decode(&tokens)
	if err != nil {
		slog.Error("decode: ", err.Error())
	}
	return
}

func (s *spotify) GetUserInfo(accessToken string) (user domain.SpotifyUser, err error) {
	u, _ := url.ParseRequestURI(domain.SpotifyBaseAPIURL + "/me")

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		slog.Error("http req: ", err.Error())
		return
	}
	req.Header.Set("Authorization", accessToken)

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("exec http req: ", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var reqErr domain.SpotifyError
		if err = json.NewDecoder(resp.Body).Decode(&reqErr); err != nil {
			slog.Error("decode resp body: ", err.Error())
			return
		}
		return user, &reqErr
	}

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		slog.Error("decode: ", err.Error())
	}
	return
}

func (s *spotify) GetUserPlaylist(accessToken, userID string) (playlists domain.SpotifyPlaylists, err error) {
	u, _ := url.ParseRequestURI(fmt.Sprintf("%v/users/%v/playlists", domain.SpotifyBaseAPIURL, userID))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		slog.Error("http req: ", err.Error())
		return
	}
	req.Header.Set("Authorization", accessToken)

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("exec http req: ", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var reqErr domain.SpotifyError
		if err = json.NewDecoder(resp.Body).Decode(&reqErr); err != nil {
			slog.Error("decode resp body: ", err.Error())
			return
		}
		return playlists, &reqErr
	}

	err = json.NewDecoder(resp.Body).Decode(&playlists)
	if err != nil {
		slog.Error("decode: ", err.Error())
	}
	return
}
