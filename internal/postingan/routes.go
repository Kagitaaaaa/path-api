package postingan

import "github.com/labstack/echo/v5"

func (h *Handler) Register(e *echo.Echo) {
	postingan := e.Group("/postingan")

	postingan.GET("", h.List)
}
