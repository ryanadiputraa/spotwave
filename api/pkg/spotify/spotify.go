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
)

type SpotifyUtil interface {
	Login(ctx *fiber.Ctx, clientID, redirectURI, state string) error
	Callback(clientID, clientSecret, code, redirectURI string) (domain.SpotifyAccessTokens, error)
}

type spotify struct{}
type AccessTokenReqErrResponse struct {
	ErrorCode        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e *AccessTokenReqErrResponse) Error() string {
	return fmt.Sprintf("Error %v : %v", e.ErrorCode, e.ErrorDescription)
}

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
		return
	}
	req.Header.Set("Authorization", "basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", clientID, clientSecret))))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var reqErr AccessTokenReqErrResponse
		if err = json.NewDecoder(resp.Body).Decode(&reqErr); err != nil {
			return
		}
		return tokens, &reqErr
	}

	err = json.NewDecoder(resp.Body).Decode(&tokens)
	return
}
