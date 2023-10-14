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
	group.Get("/callback", c.Callback)
}

func (c *controller) Login(ctx *fiber.Ctx) error {
	clientID := c.config.ClientID
	state := c.config.State
	redirectURI, _ := url.ParseRequestURI(c.config.RedirectURI)
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
