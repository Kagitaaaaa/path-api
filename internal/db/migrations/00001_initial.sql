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
	"mata_pelajaran" TEXT,
	"pembuat" INTEGER NOT NULL,
	"kode_kelas" VARCHAR(6) NOT NULL,
	"dibuat" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "anggota_kelas" (
	"id" SERIAL PRIMARY KEY,
	"pengguna_id" INTEGER NOT NULL,
	"kelas_id" INTEGER NOT NULL,
	"role" ROLE_KELAS NOT NULL,
	"dibuat" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "postingan" (
	"id" SERIAL PRIMARY KEY,
	"judul" TEXT NOT NULL,
	"deskripsi" TEXT NOT NULL,
	"lampiran" TEXT ARRAY,
	"tipe" TIPE_POST NOT NULL,
	"kelas_id" INTEGER NOT NULL,
	"dibuat" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "pengumpulan_tugas" (
	"id" SERIAL PRIMARY KEY,
	"postingan_id" INTEGER NOT NULL,
	"lampiran" TEXT ARRAY,
	"pengguna_id" INTEGER NOT NULL,
	"dibuat" TIMESTAMP NOT NULL DEFAULT now()
);

ALTER TABLE "pengumpulan_tugas"
ADD FOREIGN KEY("postingan_id") REFERENCES "postingan"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "kelas"
ADD FOREIGN KEY("pembuat") REFERENCES "pengguna"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "anggota_kelas"
ADD FOREIGN KEY("pengguna_id") REFERENCES "pengguna"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "anggota_kelas"
ADD FOREIGN KEY("kelas_id") REFERENCES "kelas"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "postingan"
ADD FOREIGN KEY("kelas_id") REFERENCES "kelas"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "pengumpulan_tugas"
ADD FOREIGN KEY("pengguna_id") REFERENCES "pengguna"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;

-- +goose Down
DROP TABLE pengumpulan_tugas;
DROP TABLE postingan;
DROP TABLE anggota_kelas;
DROP TABLE kelas;
DROP TABLE pengguna;
DROP TYPE tipe_post;
DROP TYPE role_kelas;
