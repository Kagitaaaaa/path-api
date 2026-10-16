package kelas

import (
	"log"
	"net/http"
	"path-api/internal/db"
	"path-api/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v5"
)

func (h *Handler) List(c *echo.Context) error {
	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}

	kelas, err := h.db.ListKelas(c.Request().Context(), int32(claims.UserID))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized, "tidak ada kelas")
	}

	return c.JSON(http.StatusOK, kelas)
}

func (h *Handler) Create(c *echo.Context) error {
	var req CreateKelasRequest

	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(req)
	if err != nil {
		return utils.ValidationErrorHandler(c, err)
	}

	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}

	var tingkatParam pgtype.Text
	if req.Tingkat != "" {
		tingkatParam = pgtype.Text{String: req.Tingkat, Valid: true}
	}

	kodeKelas := utils.GenerateHashCode()

	err = h.db.CreateKelas(c.Request().Context(), db.CreateKelasParams{
		Nama:      req.Nama,
		Tingkat:   tingkatParam,
		Pembuat:   int32(claims.UserID),
		KodeKelas: kodeKelas,
	})

	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal membuat kelas")
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Kelas berhasil dibuat",
	})
}

func (h *Handler) Join(c *echo.Context) error {
	var req JoinDanKeluarKelasRequest

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return utils.ValidationErrorHandler(c, err)
	}

	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}

	kelasID, err := h.db.GetKelasIDByKodeKelas(c.Request().Context(), req.KodeKelas)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "kelas tidak ditemukan")
	}

	exists, err := h.db.CheckAnggotaKelas(c.Request().Context(), db.CheckAnggotaKelasParams{
		PenggunaID: int32(claims.UserID),
		KelasID:    kelasID,
	})
	if exists != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "anda sudah bergabung di kelas ini")
	}

	err = h.db.JoinKelas(c.Request().Context(), db.JoinKelasParams{
		PenggunaID: int32(claims.UserID),
		KelasID:    kelasID,
	})
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "gagal join kelas")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Berhasil bergabung ke dalam kelas",
	})
}

func (h *Handler) Leave(c *echo.Context) error {
	var req JoinDanKeluarKelasRequest
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return utils.ValidationErrorHandler(c, err)
	}

	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}

	kelasID, err := h.db.GetKelasIDByKodeKelas(c.Request().Context(), req.KodeKelas)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "kelas tidak ditemukan")
	}

	err = h.db.LeaveKelas(c.Request().Context(), db.LeaveKelasParams{
		KelasID:    kelasID,
		PenggunaID: int32(claims.UserID),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal keluar dari kelas")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Berhasil keluar dari kelas",
	})
}
