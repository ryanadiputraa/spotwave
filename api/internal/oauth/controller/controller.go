package controller

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/config"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
	"github.com/sagikazarmark/slog-shim"
)

type controller struct {
	config *config.Config
}

func NewOauthController(group fiber.Router, config *config.Config) {
	c := controller{config: config}
	group.Get("/login", c.Login)
}

func (c *controller) Login(ctx *fiber.Ctx) error {
	clientID := c.config.Spotify.ClientID
	state := c.config.Spotify.State
	redirectURI, _ := url.ParseRequestURI(c.config.Spotify.RedirectURI)
	slog.Info(redirectURI.String())

	resource := "/authorize"
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", clientID)
	params.Add("scope", "user-read-private user-read-email")
	params.Add("redirect_uri", redirectURI.String())
	params.Add("state", state)

	u, _ := url.ParseRequestURI(domain.SpotifyAccountAPIURL)
	u.Path = resource
	u.RawQuery = params.Encode()

	return ctx.Redirect(u.String(), http.StatusTemporaryRedirect)
}
