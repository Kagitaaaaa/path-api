package kelas

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"path-api/internal/db"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v5"
)

func (h *Handler) List(c *echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "tidak ada.")
}

func generateHashKode(namaKelas string, pembuatID int32) string {
	dataUnik := fmt.Sprintf("%s-%d-%d", namaKelas, pembuatID, time.Now().UnixNano())

	hash := sha256.Sum256([]byte(dataUnik))

	hashString := hex.EncodeToString(hash[:])

	if len(hashString) >= 6 {
		return strings.ToUpper(hashString[:6])
	}

	return "CLS123"
}

func BuatKelasLogic(ctx context.Context, queries *db.Queries, req CreateKelasRequest, pembuatID int32) error {
	kodeKelas := req.KodeKelas
	if kodeKelas == "" {
		kodeKelas = generateHashKode(req.Nama, pembuatID)
	}

	err := queries.CreateKelas(ctx, db.CreateKelasParams{
		Nama:          req.Nama,
		Tingkat:       pgtype.Text{String: req.Tingkat, Valid: true},
		MataPelajaran: pgtype.Text{String: req.MataPelajaran, Valid: true},
		Pembuat:       pembuatID,
		KodeKelas:     kodeKelas,
	})

	return err
}

func JoinKelasLogic(ctx context.Context, queries *db.Queries, userID int32, kodeKelas string) error {
	kelas, err := queries.GetKelasByKode(ctx, kodeKelas)
	if err != nil {
		return errors.New("kode kelas tidak ditemukan")
	}

	sudahGabung, err := queries.CekAnggotaKelas(ctx, db.CekAnggotaKelasParams{
		PenggunaID: userID,
		KelasID:    kelas.ID,
	})
	if err == nil && sudahGabung {
		return errors.New("anda sudah bergabung di kelas ini")
	}

	err = queries.JoinKelas(ctx, db.JoinKelasParams{
		PenggunaID: userID,
		KelasID:    kelas.ID,
		Role:       "murid",
	})

	return err
}

func LeaveKelasLogic(ctx context.Context, queries *db.Queries, userID int32, kodeKelas string) error {
	// 1. Cari dulu ID kelas berdasarkan kode_kelas
	kelas, err := queries.GetKelasByKode(ctx, kodeKelas)
	if err != nil {
		return errors.New("kelas tidak ditemukan")
	}

	// 2. Hapus data keanggotaan dari tabel anggota_kelas
	err = queries.LeaveKelas(ctx, db.LeaveKelasParams{
		KelasID:    kelas.ID,
		PenggunaID: userID,
	})

	return err
}
