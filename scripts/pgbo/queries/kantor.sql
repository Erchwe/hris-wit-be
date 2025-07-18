-- name: CreateKantor :one
INSERT INTO kantor (
    kantor_id,
    nama_kantor,
    kota,
    alamat,
    nomor_telp,
    status,
    created_at,
    created_by
) VALUES (
             @kantor_id,
             @nama_kantor,
             @kota,
             @alamat,
             @nomor_telp,
             @status::status_enum,
             (now() at time zone 'UTC')::TIMESTAMP,
             @created_by
         )
    RETURNING *;

-- name: GetKantorByID :one
SELECT *
FROM kantor
WHERE kantor_id = @kantor_id
  AND deleted_at IS NULL;

-- name: ListKantor :many
SELECT *
FROM kantor
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateKantor :one
UPDATE kantor
SET
    nama_kantor = @nama_kantor,
    kota = @kota,
    alamat = @alamat,
    nomor_telp = @nomor_telp,
    status = @status::status_enum,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE kantor_id = @kantor_id
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteKantor :exec
UPDATE kantor
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE kantor_id = @kantor_id;

-- name: RestoreKantor :exec
UPDATE kantor
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE kantor_id = @kantor_id;
