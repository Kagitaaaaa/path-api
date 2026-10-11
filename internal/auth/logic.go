package auth

import (
	"log"
	"net/http"
	"path-api/internal/db"
	"path-api/internal/utils"

	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal memproses password")
	}

	exists, err := h.db.GetEmailPengguna(c.Request().Context(), req.Email)
	if exists != 0 {
		log.Println("email sudah terdaftarkan:", exists)
		return echo.NewHTTPError(http.StatusBadRequest, "email sudah terdaftarkan!")
	}

	userID, err := h.db.CreatePengguna(c.Request().Context(), db.CreatePenggunaParams{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		log.Println("Error insert pengguna:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal mendaftarkan pengguna")
	}

	token, err := utils.NewToken(int(userID))
	if err != nil {
		log.Println("Error membuat token:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal memproses register")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
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

	pengguna, err := h.db.GetPengguna(c.Request().Context(), req.Email)
	if err != nil {
		log.Println("username atau password salah", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "username atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(pengguna.Password), []byte(req.Password))
	if err != nil {
		log.Println("Error mengecek password:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "username atau password salah")
	}

	token, err := utils.NewToken(int(pengguna.ID))
	if err != nil {
		log.Println("Error membuat token:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal memproses register")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *Handler) Logout(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logout Berhasil",
	})
}