package pengguna

import (
	"path-api/internal/db"
)

type PenggunaHandler struct {
	db *db.Queries
}

func New(db *db.Queries) *PenggunaHandler {
	return &PenggunaHandler{db}
}
