package auth

import "github.com/labstack/echo/v5"

func (h *Handler) Routes(e *echo.Echo) {
	e.GET("/login", h.Login)
	e.GET("/register", h.Register)
}
