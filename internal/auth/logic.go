package auth

import (
	"net/http"
	"path-api/internal/utils"

	"github.com/labstack/echo/v5"
)

func (h *Handler) Register(c *echo.Context) error {
	req := new(RegisterRequest)

	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(req)
	if err != nil {
		return utils.ValidationErrorHandler(c, err)
	}

	return echo.NewHTTPError(200, "OK")
}

func (h *Handler) Login(c *echo.Context) error {
	req := new(LoginRequest)

	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(req)
	if err != nil {
		return utils.ValidationErrorHandler(c, err)
	}

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return echo.NewHTTPError(200, "OK")
}
