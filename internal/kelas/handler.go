package kelas

import (
	"path-api/internal/db"
)

type KelasHandler struct {
	db *db.Queries
}

func New(db *db.Queries) *KelasHandler {
	return &KelasHandler{db}
}
