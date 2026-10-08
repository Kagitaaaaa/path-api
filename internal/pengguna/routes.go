package pengguna

import "github.com/labstack/echo/v5"

func (h *Handler) Register(e *echo.Echo) {
	pengguna := e.Group("/pengguna")

	pengguna.GET("", h.List)
	pengguna.GET("/:id", h.Get)
}
