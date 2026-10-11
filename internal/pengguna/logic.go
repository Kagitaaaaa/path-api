package pengguna

import (
	"log"
	"net/http"
	"path-api/internal/db"
	"path-api/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Get(c *echo.Context) error {
	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}

	pengguna, err := h.db.GetPenggunaByID(c.Request().Context(), int32(claims.UserID))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, "Data pengguna tidak ditemukan")
	}

	return c.JSON(http.StatusOK, pengguna)
}

func (h *Handler) UpdateEmail(c *echo.Context) error {
	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}

	req := new(UpdateEmailRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return utils.ValidationErrorHandler(c, err)
	}

	updatedUser, err := h.db.UpdateEmailPengguna(c.Request().Context(), db.UpdateEmailPenggunaParams{
		ID:    int32(claims.UserID),
		Email: req.Email,
	})
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal memperbarui email")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Email berhasil diperbarui",
		"data":    updatedUser,
	})
}

func (h *Handler) UpdatePhone(c *echo.Context) error {
	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}

	req := new(UpdatePhoneRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return utils.ValidationErrorHandler(c, err)
	}

	phonePg := pgtype.Text{String: req.Phone, Valid: true}

	updatedUser, err := h.db.UpdatePhonePengguna(c.Request().Context(), db.UpdatePhonePenggunaParams{
		ID:    int32(claims.UserID),
		Phone: phonePg,
	})
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal memperbarui nomor HP")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Nomor HP berhasil diperbarui",
		"data":    updatedUser,
	})
}

func (h *Handler) UpdateProfilePicture(c *echo.Context) error {
	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}

	req := new(UpdateProfilePictureRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return utils.ValidationErrorHandler(c, err)
	}

	ppPg := pgtype.Text{String: req.ProfilePicture, Valid: true}

	updatedUser, err := h.db.UpdateProfilePicturePengguna(c.Request().Context(), db.UpdateProfilePicturePenggunaParams{
		ID:             int32(claims.UserID),
		ProfilePicture: ppPg,
	})
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal memperbarui foto profil")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Foto profil berhasil diperbarui",
		"data":    updatedUser,
	})
}

func (h *Handler) ResetPassword(c *echo.Context) error {
	claims, err := utils.GetClaims(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Tidak dapat mengakses data klaim")
	}

	req := new(ResetPasswordRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return utils.ValidationErrorHandler(c, err)
	}

	user, err := h.db.GetPenggunaByID(c.Request().Context(), int32(claims.UserID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Pengguna tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Password tidak sesuai dengan Password lama")
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal memproses password baru")
	}

	err = h.db.UpdatePassword(c.Request().Context(), db.UpdatePasswordParams{
		ID:       user.ID,
		Password: string(hashedNewPassword),
	})
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal mereset password")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Reset password berhasil",
	})
}