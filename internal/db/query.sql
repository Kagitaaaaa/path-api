---- AUTH ----
-- name: GetPasswordPenggunaByEmail :one
SELECT id, password FROM pengguna WHERE email = $1;

-- name: GetPasswordPenggunaByID :one
SELECT password FROM pengguna WHERE id = $1;


---- PENGGUNA ----
-- name: CreatePengguna :one
INSERT INTO pengguna (username, email, password)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetPengguna :one
SELECT id, username, email, phone, profile_picture FROM pengguna WHERE id = $1 LIMIT 1;

-- name: UpdatePasswordPengguna :exec
UPDATE pengguna SET password = $2 WHERE id = $1;

-- name: UpdateEmailPengguna :one
UPDATE pengguna SET email = $2 WHERE id = $1 RETURNING id, username, email, phone, profile_picture;

-- name: UpdatePhonePengguna :one
UPDATE pengguna SET phone = $2 WHERE id = $1 RETURNING id, username, email, phone, profile_picture;

-- name: UpdateProfilePicturePengguna :one
UPDATE pengguna SET profile_picture = $2 WHERE id = $1 RETURNING id, username, email, phone, profile_picture;


---- KELAS ----
-- name: CreateKelas :exec
INSERT INTO kelas (nama, tingkat, pembuat, kode_kelas)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: ListKelas :many
SELECT kelas.* FROM anggota_kelas
JOIN kelas ON kelas_id = kelas.id
WHERE pengguna_id = $1;

-- name: GetKelas :one
SELECT * FROM kelas WHERE kode_kelas = $1;

-- name: JoinKelas :exec
INSERT INTO anggota_kelas (pengguna_id, kelas_id, role)
VALUES ($1, $2, $3);

-- name: GetKelasIDByKodeKelas :one
SELECT id FROM kelas WHERE kode_kelas = $1;

-- name: UpdateRoleKelas :exec
UPDATE anggota_kelas
SET role = $2
WHERE id = $1;

-- name: ListAnggotaKelas :one
SELECT * FROM anggota_kelas
WHERE kelas_id = $1;

-- name: CheckAnggotaKelas :one
SELECT 1 FROM anggota_kelas
WHERE pengguna_id = $1 AND kelas_id = $2;

-- name: LeaveKelas :exec
DELETE FROM anggota_kelas
WHERE kelas_id = $1 AND pengguna_id = $2;

---- POSTINGAN ----
-- name: ListPostingan :many
SELECT p.id, p.judul, p.deskripsi, p.lampiran, p.tipe, p.pengguna_id, p.kelas_id, p.dibuat
FROM postingan p
JOIN kelas k ON p.kelas_id = k.id
WHERE p.tipe = $1
  AND p.pengguna_id = $2
  AND p.kelas_id = $3
  AND p.mata_pelajaran = $4;

-- name: CreatePostingan :one
INSERT INTO postingan (judul, deskripsi, lampiran, tipe, pengguna_id, kelas_id, mata_pelajaran)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdatePostingan :one
UPDATE postingan
SET judul = $2, deskripsi = $3, lampiran = $4, tipe = $5
WHERE id = $1
RETURNING *;
