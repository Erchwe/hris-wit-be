-- name: CreatePelamarReferensi :one
INSERT INTO pelamar_referensi (
    id_referensi, id_pelamar, nama, nama_perusahaan, jabatan, no_telp_perusahaan,
    created_at, created_by
) VALUES (
    @id_referensi, @id_pelamar, @nama, @nama_perusahaan, @jabatan, @no_telp_perusahaan,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPelamarReferensiByID :one
SELECT * FROM pelamar_referensi
WHERE id_referensi = @id_referensi
  AND deleted_at IS NULL;

-- name: ListPelamarReferensi :many
SELECT * FROM pelamar_referensi
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdatePelamarReferensi :one
UPDATE pelamar_referensi
SET
    nama = @nama,
    nama_perusahaan = @nama_perusahaan,
    jabatan = @jabatan,
    no_telp_perusahaan = @no_telp_perusahaan,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_referensi = @id_referensi
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePelamarReferensi :exec
UPDATE pelamar_referensi
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_referensi = @id_referensi
  AND deleted_at IS NULL;

-- name: RestorePelamarReferensi :exec
UPDATE pelamar_referensi
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_referensi = @id_referensi;
