package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanadiputraa/spotwave/api/config"
	"github.com/ryanadiputraa/spotwave/api/internal/domain"
	"github.com/ryanadiputraa/spotwave/api/internal/spotify"
	rapidapi "github.com/ryanadiputraa/spotwave/api/pkg/rapid-api"
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
	group.Get("/playlists", c.GetUserPlaylists)
	group.Get("/playlists/tracks", c.GetPlaylistTracks)
	group.Get("/tracks/download", c.DownloadSpotifyTrack)
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

func (c *controller) GetUserPlaylists(ctx *fiber.Ctx) error {
	context := ctx.Context()
	headers := ctx.GetReqHeaders()
	accessToken := headers["Authorization"]

	user, err := c.service.GetUserPlaylists(context, accessToken)
	if err != nil {
		if spotifyErr, ok := err.(*domain.SpotifyError); ok {
			slog.Warn("spotify error: " + spotifyErr.ErrorDetail.Message)
			return ctx.Status(spotifyErr.ErrorDetail.Status).JSON(fiber.Map{
				"error":   domain.ErrBadRequest,
				"message": spotifyErr.ErrorDetail.Message,
			})
		}
		slog.Warn("fail to get user playlists: " + err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   domain.ErrBadRequest,
			"message": "fail to get user playlists",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "fetch spotify user playlists",
		"data":    user,
	})
}

func (c *controller) GetPlaylistTracks(ctx *fiber.Ctx) error {
	context := ctx.Context()
	headers := ctx.GetReqHeaders()
	accessToken := headers["Authorization"]
	m := ctx.Queries()
	playlistID := m["playlist_id"]

	tracks, err := c.service.GetPlaylistTracks(context, accessToken, playlistID)
	if err != nil {
		if spotifyErr, ok := err.(*domain.SpotifyError); ok {
			slog.Warn("spotify error: " + spotifyErr.ErrorDetail.Message)
			return ctx.Status(spotifyErr.ErrorDetail.Status).JSON(fiber.Map{
				"error":   domain.ErrBadRequest,
				"message": spotifyErr.ErrorDetail.Message,
			})
		}
		slog.Warn("fail to get playlist tracks: " + err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   domain.ErrBadRequest,
			"message": "fail to get playlist tracks",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "fetch spotify playlist tracks",
		"data":    tracks,
	})
}

func (c *controller) DownloadSpotifyTrack(ctx *fiber.Ctx) error {
	context := ctx.Context()
	m := ctx.Queries()
	artists := m["artists"]
	title := m["title"]

	link, err := c.service.DownloadTrack(context, artists, title)
	slog.Error("err: ", err)
	if err != nil {
		if rapidAPIErr, ok := err.(*rapidapi.ErrorResponse); ok {
			slog.Warn("rapid api error: " + rapidAPIErr.ErrorMessage)
			return ctx.Status(rapidAPIErr.Code).JSON(fiber.Map{
				"error":   domain.ErrBadRequest,
				"message": rapidAPIErr.Message,
			})
		}
		slog.Warn("fail to download track: " + err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   domain.ErrBadRequest,
			"message": "fail to download track",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "track downloaded",
		"data": fiber.Map{
			"link": link,
		},
	})
}
