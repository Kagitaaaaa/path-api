package pengguna

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *Handler) Get(c *echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "tidak ada.")
}
