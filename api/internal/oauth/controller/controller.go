package controller

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/config"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
	"github.com/ryanadiputraa/spotwave/api/internal/oauth"
	"github.com/ryanadiputraa/spotwave/api/pkg/spotify"
	"github.com/sagikazarmark/slog-shim"
)

type controller struct {
	config      *config.Config
	service     oauth.Usecase
	spotifyUtil spotify.SpotifyUtil
}

func NewController(group fiber.Router, config *config.Config, service oauth.Usecase, spotifyUtil spotify.SpotifyUtil) {
	c := controller{config: config, service: service, spotifyUtil: spotifyUtil}
	group.Get("/login", c.Login)
	group.Get("/callback", c.Callback)
	group.Get("/refresh_token", c.RefreshToken)
}

func (c *controller) Login(ctx *fiber.Ctx) error {
	clientID := c.config.ClientID
	redirectURI, _ := url.ParseRequestURI(c.config.RedirectURI)
	state := c.config.State

	return c.spotifyUtil.Login(ctx, clientID, redirectURI.String(), state)
}

func (c *controller) Callback(ctx *fiber.Ctx) error {
	context := ctx.Context()
	m := ctx.Queries()
	u, _ := url.ParseRequestURI(c.config.WebURL)
	u.Path = "/auth"
	params := url.Values{}

	state := m["state"]
	if state != c.config.State {
		slog.Warn("oath callback: invalid state param")
		params.Set("err", domain.ErrForbidden)
		u.RawQuery = params.Encode()
		return ctx.Redirect(u.String(), http.StatusTemporaryRedirect)
	}

	error := m["error"]
	if error != "" {
		slog.Warn("oath callback: ", error)
		params.Set("err", error)
		u.RawQuery = params.Encode()
		return ctx.Redirect(u.String(), http.StatusTemporaryRedirect)
	}

	code := m["code"]
	tokens, err := c.service.Callback(context, code)
	if err != nil {
		if spotifyErr, ok := err.(*domain.SpotifyOauthError); ok {
			slog.Warn("spotify error: ", spotifyErr.ErrorDescription)
			params.Set("err", spotifyErr.ErrorCode)
			u.RawQuery = params.Encode()
			return ctx.Redirect(u.String(), http.StatusTemporaryRedirect)
		}
		slog.Warn("oath get spotify access token: ", err)
		params.Set("err", domain.ErrBadRequest)
		u.RawQuery = params.Encode()
		return ctx.Redirect(u.String(), http.StatusTemporaryRedirect)
	}

	expires := strconv.Itoa(tokens.ExpiresIn)
	params.Add("access_token", tokens.AccessToken)
	params.Add("expires_in", expires)
	params.Add("refresh_token", tokens.RefreshToken)
	u.RawQuery = params.Encode()

	return ctx.Redirect(u.String(), http.StatusTemporaryRedirect)
}

func (c *controller) RefreshToken(ctx *fiber.Ctx) error {
	context := ctx.Context()
	m := ctx.Queries()

	refreshToken := m["token"]
	if refreshToken == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   domain.ErrBadRequest,
			"message": "missing 'token' param",
		})
	}

	tokens, err := c.service.RefreshToken(context, refreshToken)
	if err != nil {
		if spotifyErr, ok := err.(*domain.SpotifyOauthError); ok {
			slog.Warn("spotify error: ", spotifyErr.ErrorDescription)
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error":   spotifyErr.ErrorCode,
				"message": spotifyErr.ErrorDescription,
			})
		}
		slog.Warn("fail to refresh token: ", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   http.StatusBadRequest,
			"message": "fail to refresh access token",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "refresh token success",
		"data":    tokens,
	})
}
