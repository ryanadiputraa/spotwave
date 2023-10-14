package spotify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
)

type SpotifyUtil interface {
	Login(ctx *fiber.Ctx, clientID, redirectURI, state string) error
	Callback(clientID, clientSecret, code, redirectURI string) (domain.SpotifyAccessTokens, error)
}

type spotify struct{}

type AccessTokenDTO struct {
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
	GrantType   string `json:"grant_type"`
}

func NewSpotifyUtil() SpotifyUtil {
	return &spotify{}
}

func (s *spotify) Login(ctx *fiber.Ctx, clientID, redirectURI, state string) error {
	resource := "/authorize"
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", clientID)
	params.Add("scope", "user-read-private user-read-email")
	params.Add("redirect_uri", redirectURI)
	params.Add("state", state)
	params.Add("show_dialog", "true")

	u, _ := url.ParseRequestURI(domain.SpotifyAccountAPIURL)
	u.Path = resource
	u.RawQuery = params.Encode()

	return ctx.Redirect(u.String(), http.StatusTemporaryRedirect)
}

func (s *spotify) Callback(clientID, clientSecret, code, redirectURI string) (tokens domain.SpotifyAccessTokens, err error) {
	recourse := "/api/token"
	u, _ := url.ParseRequestURI(domain.SpotifyAccountAPIURL)
	u.Path = recourse

	dto := AccessTokenDTO{
		Code:        code,
		RedirectURI: redirectURI,
		GrantType:   "authorization_code",
	}
	body, err := json.Marshal(dto)
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return
	}
	req.Header.Set("Authorization", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", clientID, clientSecret))))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&tokens)
	return
}
