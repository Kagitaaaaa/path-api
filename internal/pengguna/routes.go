package pengguna

import (
	// "path-api/internal/utils"

	// "github.com/golang-jwt/jwt/v5"
	// echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func (h *Handler) Routes(e *echo.Echo) {
	// customJWTConfig := echojwt.Config{
	// 	NewClaimsFunc: func(c *echo.Context) jwt.Claims {
	// 		return new(utils.JWTCustomClaims) 
	// 	},
	// 	SigningKey: []byte(utils.GetSecret()),
	// }
	// pengguna := e.Group("/pengguna", echojwt.WithConfig(customJWTConfig))
	pengguna := e.Group("/pengguna")

	pengguna.GET("", h.Get)
	pengguna.PATCH("/email", h.UpdateEmail)
	pengguna.PATCH("/phone", h.UpdatePhone)
	pengguna.PATCH("/profile-picture", h.UpdateProfilePicture)
	pengguna.PATCH("/reset-password", h.ResetPassword)
}