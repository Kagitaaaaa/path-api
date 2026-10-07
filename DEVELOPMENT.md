# Panduan Development PATH API

## Daftar Isi
- [Prerequisites](#prerequisites)
- [Membuat query SQL](#membuat-query-sql)
- [Routing](#routing)
- [Membuat Implementasi Route](#membuat-implementasi-route)
- [Mengakses query SQL di `logic.go`](#mengakses-query-sql-di-logicgo)
- [Catatan Tambahan](#catatan-tambahan)

## Prerequisites
1. Go version >= 1.26
2. Postgres >= 1.18
3. sqlc 
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
atau kunjungi [instalasi sqlc](https://docs.sqlc.dev/en/latest/overview/install.html).

## Membuat query SQL
1. Masuk pada file `internal/db/query.sql`
2. Buat anotasi sebelum query dengan
```sql
-- name: NAMA :TIPE
```
jenis TIPE ada 3, yakni:
- `:exec` untuk kode yang hanya eksekusi dan tidak menghasilkan apapun. \
contoh:
```sql
-- name: DeletePengguna :exec
DELETE FROM pengguna WHERE id = $1;

-- name: UpdateAlamatPengguna :exec
UPDATE pengguna SET alamat = 'Kentingan' WHERE id = $1;
```

- `:one` untuk kode yang hanya menghasilkan satu row. \
contoh:
```sql
-- name: GetPengguna :one
SELECT * FROM pengguna WHERE id = $1;

-- name: UpdateAlamatPenggunaWithReturning :one
UPDATE pengguna SET alamat = 'Kentingan' WHERE id = $1 RETURNING *;
```

- `:many` untuk kode yang menghasilkan lebih dari satu row. \
contoh:
```sql
-- name: ListPengguna :many
SELECT * FROM pengguna;
```

3. Setelah modifikasi `query.sql`. Jalankan 
```bash
sqlc generate
```

Selengkapnya dapat cek [dokumentasi sqlc](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)


## Routing
1. Routing pada PATH API terpisah sesuai dengan bidang. Untuk memberikan routing terkait dengan pengguna maka masuk pada file `internal/pengguna/routes.go`
2. Beri route pada fungsi `Register()`. \
Contoh:
```go
func (h *PenggunaHandler) Register(e *echo.Echo) {
  //     path        logic
  e.GET("/pengguna", h.List)
  e.POST("/pengguna", h.Create)
  e.GET("/pengguna/:id", h.Get)
  e.DELETE("/pengguna/:id", h.Delete)
}
```
`h.List`, `h.Create`, `h.Get`, `h.Delete` adalah contoh fungsi yang kita buat, cara membuatnya lanjut pada bagian [Membuat Implementasi Route](#membuat-implementasi-route)

## Membuat Implementasi Route
1. Masih sama pada konteks pengguna, masuk pada `internal/pengguna/logic.go`
2. Buat fungsi dengan boilerplate seperti ini
```go
func (h *PenggunaHandler) List(c *echo.Context) error {
  // ...
}
```

3. Untuk menghasilkan response, gunakan variabel `c *echo.Context`. \
Contoh:
```go
func (h *PenggunaHandler) List(c *echo.Context) error {
  // menghasilkan response berbentuk string
  return c.String(http.StatusOK, "Hello, World!")
  
  // atau
  
  // menghasilkan response berbentuk json
  return c.JSON(http.StatusOK, map[string]string{
    "nama": "John Doe",
    "alamat": "Solo",
  }) 
}

```

## Mengakses query SQL di `logic.go`
Setelah generate query dari section [Membuat query SQL](#membuat-query-sql). Seperti contoh:
```sql
-- name: ListPengguna :many
SELECT * FROM pengguna;
```

lalu menjalankan:

```bash
sqlc generate
```

Kita bisa langsung mengakses fungsi yang digenerate sqlc di `logic.go`dengan variabel `h.db`.

Contoh:
```go
func (h *PenggunaHandler) List(c *echo.Context) error {
	pengguna, err := h.db.ListPengguna(c.Request().Context())
	if err != nil {
    // ...
	}

	return c.JSON(http.StatusBadRequest, pengguna)
}

```

## Catatan Tambahan
untuk fungsionalitas Framework Echo yang lain, selengkapnya dapat mengecek dokumentasi [Echo](https://echo.labstack.com/guide/quickstart/)

> [!WARNING]
> Agar pola response lebih teratur dan terprediksi pada pemanggilan dari frontend. JANGAN MEMBERIKAN RESPONSE BERUPA STRING\
> \
> Seperti contoh: 
> ```go
> return c.String(http.StatusOK, "JANGAN LAKUKAN INI")
> ```
> \
> **SELALU GUNAKAN JSON!**
> Contoh:
> ```go
> return c.JSON(http.StatusOK, "{\"message\": \"OK\"}")
> 
> // atau semisal ingin mengembalikan error dapat menggunakan error handling bawaan Echo
> return echo.NewHTTPError(http.StatusBadRequest, "ini message error")
> ```
