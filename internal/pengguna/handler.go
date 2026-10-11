package pengguna

import (
	"path-api/internal/db"
)

type Handler struct {
	db *db.Queries
}

func New(db *db.Queries) *Handler {
	return &Handler{db}
}

type UpdateEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type UpdatePhoneRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type UpdateProfilePictureRequest struct {
	ProfilePicture string `json:"profile_picture" validate:"required"`
}

type ResetPasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}