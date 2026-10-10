-- name: CreatePengguna :one
INSERT INTO pengguna (username, email, password)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetEmailPengguna :one
SELECT 1 FROM pengguna WHERE email = $1;

-- name: GetPengguna :one
SELECT id, password FROM pengguna WHERE email = $1;
