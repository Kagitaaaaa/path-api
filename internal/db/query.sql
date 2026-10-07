-- name: ListPengguna :many
select * from pengguna;

-- name: GetPengguna :one
select * from pengguna where id = $1;
