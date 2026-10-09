package pengguna

import (
	"log"

	"github.com/labstack/echo/v5"
)

func (h *Handler) List(c *echo.Context) error {
	pengguna, err := h.db.ListPengguna(c.Request().Context())
	if err != nil {
		log.Println(err)
		return c.String(500, "gagal mendapatkan data pengguna")
	}

	return c.JSON(200, pengguna)
}

func (h *Handler) Get(c *echo.Context) error {
	// 	idStr := c.Param("id")
	//
	// 	id, _ := strconv.Atoi(idStr)
	//
	// 	pengguna, err := h.db.GetPengguna(c.Request().Context(), int32(id))
	// 	if err != nil {
	// 		log.Println(err)
	// 		return c.String(500, "gagal mendapatkan data pengguna")
	// 	}

	return echo.NewHTTPError(400, "tidak ada.")
}

func (h *Handler) Create(c *echo.Context) error {
	return echo.NewHTTPError(400, "Bad Request")
}
