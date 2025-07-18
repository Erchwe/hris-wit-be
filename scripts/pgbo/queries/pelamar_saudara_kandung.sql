-- name: CreatePelamarSaudaraKandung :one
INSERT INTO pelamar_saudara_kandung (
    id_saudara, id_pelamar, nama, jenis_kelamin, tempat_lahir,
    pendidikan_pekerjaan, tgl_lahir, created_at, created_by
) VALUES (
    @id_saudara, @id_pelamar, @nama, @jenis_kelamin, @tempat_lahir,
    @pendidikan_pekerjaan, @tgl_lahir, NOW(), @created_by
) RETURNING *;

-- name: GetPelamarSaudaraKandungByID :one
SELECT * FROM pelamar_saudara_kandung
WHERE id_saudara = @id_saudara AND deleted_at IS NULL;

-- name: ListPelamarSaudaraKandung :many
SELECT * FROM pelamar_saudara_kandung
WHERE deleted_at IS NULL ORDER BY created_at DESC;

-- name: UpdatePelamarSaudaraKandung :one
UPDATE pelamar_saudara_kandung
SET
    nama = @nama,
    jenis_kelamin = @jenis_kelamin,
    tempat_lahir = @tempat_lahir,
    pendidikan_pekerjaan = @pendidikan_pekerjaan,
    tgl_lahir = @tgl_lahir,
    updated_at = NOW(),
    updated_by = @updated_by
WHERE id_saudara = @id_saudara AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePelamarSaudaraKandung :exec
UPDATE pelamar_saudara_kandung
SET deleted_at = NOW(), deleted_by = @deleted_by
WHERE id_saudara = @id_saudara AND deleted_at IS NULL;

-- name: RestorePelamarSaudaraKandung :exec
UPDATE pelamar_saudara_kandung
SET deleted_at = NULL, deleted_by = NULL
WHERE id_saudara = @id_saudara AND deleted_at IS NOT NULL;
