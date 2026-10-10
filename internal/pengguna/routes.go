package pengguna

import "github.com/labstack/echo/v5"

func (h *Handler) Routes(e *echo.Echo) {
	pengguna := e.Group("/pengguna")

	pengguna.GET("", h.Get)
}
