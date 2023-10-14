package controller

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/config"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
	"github.com/ryanadiputraa/spotwave/api/pkg/spotify"
	"github.com/sagikazarmark/slog-shim"
)

type controller struct {
	config      *config.Config
	spotifyUtil spotify.SpotifyUtil
}

func NewOauthController(group fiber.Router, config *config.Config, spotifyUtil spotify.SpotifyUtil) {
	c := controller{config: config, spotifyUtil: spotifyUtil}
	group.Get("/login", c.Login)
	group.Get("/callback", c.Callback)
}

func (c *controller) Login(ctx *fiber.Ctx) error {
	clientID := c.config.ClientID
	redirectURI, _ := url.ParseRequestURI(c.config.RedirectURI)
	state := c.config.State

	return c.spotifyUtil.Login(ctx, clientID, redirectURI.String(), state)
}

func (c *controller) Callback(ctx *fiber.Ctx) error {
	m := ctx.Queries()

	state := m["state"]
	if state != c.config.State {
		slog.Warn("oath callback: invalid state param")
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
			"err_code": domain.ErrForbidden,
			"message":  "invalid state param",
		})
	}

	error := m["error"]
	if error != "" {
		slog.Warn("oath callback: ", error)
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
			"err_code": error,
			"message":  "fail to login",
		})
	}

	code := m["code"]
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"code": code,
	})
}
