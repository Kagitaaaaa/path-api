package kelas

import "github.com/labstack/echo/v5"

func (h *Handler) Routes(e *echo.Group) {
	kelas := e.Group("/kelas")

	kelas.GET("", h.List)
	kelas.POST("", h.Create)
	kelas.POST("/gabung", h.Join)
	kelas.DELETE("/keluar", h.Leave)
}
