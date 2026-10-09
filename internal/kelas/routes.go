package kelas

import "github.com/labstack/echo/v5"

func (h *Handler) Routes(e *echo.Echo) {
	kelas := e.Group("/kelas")

	kelas.GET("", h.List)
}
