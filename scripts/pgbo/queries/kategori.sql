-- name: CreateKategori :one
INSERT INTO kategori (
    kategori_id,
    nama_kategori,
    created_at,
    created_by
) VALUES (
             @kategori_id,
             @nama_kategori,
             (now() at time zone 'UTC')::TIMESTAMP,
             @created_by
         )
    RETURNING *;

-- name: GetKategoriByID :one
SELECT *
FROM kategori
WHERE kategori_id = @kategori_id
  AND deleted_at IS NULL;

-- name: ListKategori :many
SELECT *
FROM kategori
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateKategori :one
UPDATE kategori
SET
    nama_kategori = @nama_kategori,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE kategori_id = @kategori_id
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteKategori :exec
UPDATE kategori
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE kategori_id = @kategori_id;

-- name: RestoreKategori :exec
UPDATE kategori
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE kategori_id = @kategori_id;
