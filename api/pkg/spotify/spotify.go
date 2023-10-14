package spotify

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
)

type SpotifyUtil interface {
	Login(ctx *fiber.Ctx, clientID, redirectURI, state string) error
}

type spotify struct{}

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
