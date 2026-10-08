-- name: ListPengguna :many
select * from pengguna;

-- name: GetPengguna :one
select id, name from pengguna where id = $1;
