package pengguna

import (
	"github.com/labstack/echo/v5"
)

func (h *PenggunaHandler) Register(e *echo.Echo) {
	e.GET("/pengguna", h.List)
	e.POST("/pengguna", h.Create)
	e.GET("/pengguna/:id", h.Get)
}
