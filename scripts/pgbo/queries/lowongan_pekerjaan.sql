-- name: CreateLowonganPekerjaan :one
INSERT INTO lowongan_pekerjaan (
    id_lowongan_pekerjaan, posisi, tgl_buka_lowongan, tgl_tutup_lowongan,
    kriteria, deskripsi, link_lowongan, created_at, created_by
) VALUES (
    @id_lowongan_pekerjaan, @posisi, @tgl_buka_lowongan, @tgl_tutup_lowongan,
    @kriteria, @deskripsi, @link_lowongan, (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetLowonganPekerjaanByID :one
SELECT * FROM lowongan_pekerjaan
WHERE id_lowongan_pekerjaan = @id_lowongan_pekerjaan
  AND deleted_at IS NULL;

-- name: ListLowonganPekerjaan :many
SELECT * FROM lowongan_pekerjaan
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateLowonganPekerjaan :one
UPDATE lowongan_pekerjaan
SET
    posisi = @posisi,
    tgl_buka_lowongan = @tgl_buka_lowongan,
    tgl_tutup_lowongan = @tgl_tutup_lowongan,
    kriteria = @kriteria,
    deskripsi = @deskripsi,
    link_lowongan = @link_lowongan,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_lowongan_pekerjaan = @id_lowongan_pekerjaan
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteLowonganPekerjaan :exec
UPDATE lowongan_pekerjaan
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_lowongan_pekerjaan = @id_lowongan_pekerjaan
  AND deleted_at IS NULL;

-- name: RestoreLowonganPekerjaan :exec
UPDATE lowongan_pekerjaan
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_lowongan_pekerjaan = @id_lowongan_pekerjaan;

-- name: GetLastLowonganPekerjaanID :one
SELECT id_lowongan_pekerjaan
FROM lowongan_pekerjaan
ORDER BY id_lowongan_pekerjaan DESC
    LIMIT 1;

