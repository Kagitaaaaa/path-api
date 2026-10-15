package kelas

import (
	"encoding/json"
	"net/http"
	"path-api/internal/db"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

type Handler struct {
	db *db.Queries
}

func New(db *db.Queries) *Handler {
	return &Handler{db}
}

type CreateKelasRequest struct {
	Nama          string `json:"nama"`
	Tingkat       string `json:"tingkat"`
	MataPelajaran string `json:"mata_pelajaran"`
	Pembuat       int32  `json:"pembuat"`
	KodeKelas     string `json:"kode_kelas"`
}

func (h *Handler) Create(c *echo.Context) error {
	var req CreateKelasRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Format input JSON tidak valid")
	}

	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak ditemukan")
	}

	claimsJSON, _ := json.Marshal(userToken.Claims)
	var claims map[string]interface{}
	json.Unmarshal(claimsJSON, &claims)

	idClaim, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Key 'user_id' tidak ditemukan di token")
	}

	req.Pembuat = int32(idClaim)
	if req.Nama == "" {
    return echo.NewHTTPError(http.StatusBadRequest, "Nama kelas wajib diisi")
	}

	err := BuatKelasLogic(c.Request().Context(), h.db, req, req.Pembuat)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal membuat kelas: "+err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Kelas berhasil dibuat",
	})
}

type JoinKelasRequest struct {
	KodeKelas string `json:"kode_kelas"`
}

func (h *Handler) Join(c *echo.Context) error {
	var req JoinKelasRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Format input JSON tidak valid")
	}

	if req.KodeKelas == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Kode kelas wajib diisi")
	}

	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak ditemukan")
	}
	
	claimsJSON, _ := json.Marshal(userToken.Claims)
	var claims map[string]interface{}
	json.Unmarshal(claimsJSON, &claims)

	idClaim, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Gagal membaca ID dari token")
	}
	userID := int32(idClaim)

	// Panggil logic join kelas
	err := JoinKelasLogic(c.Request().Context(), h.db, userID, req.KodeKelas)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Berhasil bergabung ke dalam kelas",
	})
}

func (h *Handler) Leave(c *echo.Context) error {
	var req JoinKelasRequest // Bisa menggunakan struct yang sama karena hanya butuh "kode_kelas"
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Format input JSON tidak valid")
	}

	if req.KodeKelas == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Kode kelas wajib diisi")
	}

	// Ambil ID user secara aman dari JWT
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak ditemukan")
	}
	
	claimsJSON, _ := json.Marshal(userToken.Claims)
	var claims map[string]interface{}
	json.Unmarshal(claimsJSON, &claims)

	idClaim, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Gagal membaca ID dari token")
	}
	userID := int32(idClaim)

	err := LeaveKelasLogic(c.Request().Context(), h.db, userID, req.KodeKelas)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Gagal keluar dari kelas: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Berhasil keluar dari kelas",
	})
}