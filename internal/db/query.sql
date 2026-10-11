-- name: CreatePengguna :one
INSERT INTO pengguna (username, email, password)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetEmailPengguna :one
SELECT 1 FROM pengguna WHERE email = $1;

-- name: GetPengguna :one
SELECT id, password FROM pengguna WHERE email = $1;

-- name: GetPenggunaByID :one
SELECT * FROM pengguna WHERE id = $1 LIMIT 1;

-- name: UpdatePassword :exec
UPDATE pengguna SET password = $2 WHERE id = $1;

-- name: UpdateEmailPengguna :one
UPDATE pengguna SET email = $2 WHERE id = $1 RETURNING *;

-- name: UpdatePhonePengguna :one
UPDATE pengguna SET phone = $2 WHERE id = $1 RETURNING *;

-- name: UpdateProfilePicturePengguna :one
UPDATE pengguna SET profile_picture = $2 WHERE id = $1 RETURNING *;