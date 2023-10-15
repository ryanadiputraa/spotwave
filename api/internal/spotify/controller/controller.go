package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/config"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
	"github.com/ryanadiputraa/spotwave/api/internal/spotify"
	"github.com/sagikazarmark/slog-shim"
)

type controller struct {
	config  *config.Config
	service spotify.Usecase
}

func NewController(group fiber.Router, config *config.Config, service spotify.Usecase) {
	c := controller{
		config:  config,
		service: service,
	}
	group.Get("/users", c.GetUserInfo)
}

func (c *controller) GetUserInfo(ctx *fiber.Ctx) error {
	context := ctx.Context()
	headers := ctx.GetReqHeaders()
	accessToken := headers["Authorization"]

	user, err := c.service.GetUserInfo(context, accessToken)
	if err != nil {
		if spotifyErr, ok := err.(*domain.SpotifyError); ok {
			slog.Warn("spotify error: " + spotifyErr.ErrorDetail.Message)
			return ctx.Status(spotifyErr.ErrorDetail.Status).JSON(fiber.Map{
				"error":   domain.ErrUnauthorized,
				"message": spotifyErr.ErrorDetail.Message,
			})
		}
		slog.Warn("fail to get user info: " + err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   domain.ErrBadRequest,
			"message": "fail to get user info",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "fetch spotify user info",
		"data":    user,
	})
}
