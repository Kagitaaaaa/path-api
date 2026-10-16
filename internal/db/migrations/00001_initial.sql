-- +goose Up
CREATE TYPE "role_kelas" AS ENUM (
	'pengajar',
	'murid'
);

CREATE TYPE "tipe_post" AS ENUM (
	'materi',
	'tugas'
);

CREATE TABLE IF NOT EXISTS "pengguna" (
	"id" SERIAL PRIMARY KEY,
	"username" TEXT NOT NULL,
	"email" TEXT NOT NULL UNIQUE,
	"phone" TEXT,
	"profile_picture" TEXT,
	"password" TEXT NOT NULL,
	"dibuat" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "kelas" (
	"id" SERIAL PRIMARY KEY,
	"nama" TEXT NOT NULL,
	"tingkat" TEXT,
	"pembuat" INTEGER REFERENCES pengguna (id) NOT NULL,
	"kode_kelas" VARCHAR(6) NOT NULL UNIQUE,
	"dibuat" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "anggota_kelas" (
	"id" SERIAL PRIMARY KEY,
	"pengguna_id" INTEGER REFERENCES pengguna (id) NOT NULL,
	"kelas_id" INTEGER REFERENCES kelas (id) NOT NULL,
	"role" ROLE_KELAS NOT NULL,
	"dibuat" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "postingan" (
	"id" SERIAL PRIMARY KEY,
	"judul" TEXT NOT NULL,
	"deskripsi" TEXT,
	"lampiran" TEXT ARRAY,
	"tipe" TIPE_POST NOT NULL,
	"mata_pelajaran" TEXT,
	"pengguna_id" INTEGER REFERENCES pengguna (id) NOT NULL,
	"kelas_id" INTEGER REFERENCES kelas (id) NOT NULL,
	"dibuat" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "pengumpulan_tugas" (
	"id" SERIAL PRIMARY KEY,
	"postingan_id" INTEGER REFERENCES postingan (id) NOT NULL,
	"lampiran" TEXT ARRAY,
	"dibuat" TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE pengumpulan_tugas;
DROP TABLE postingan;
DROP TABLE anggota_kelas;
DROP TABLE kelas;
DROP TABLE pengguna;
DROP TYPE tipe_post;
DROP TYPE role_kelas;
