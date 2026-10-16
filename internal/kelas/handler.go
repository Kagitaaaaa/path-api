package kelas

import (
	"path-api/internal/db"
)

type Handler struct {
	db *db.Queries
}

func New(db *db.Queries) *Handler {
	return &Handler{db}
}

type CreateKelasRequest struct {
	Nama    string `json:"nama" validate:"required"`
	Tingkat string `json:"tingkat"`
}

type JoinDanKeluarKelasRequest struct {
	KodeKelas string `param:"kode_kelas" validate:"required"`
}
