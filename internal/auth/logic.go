package auth

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *Handler) Login(c *echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "tidak ada.")
}

func (h *Handler) Register(c *echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "tidak ada.")
}
