-- name: CreateRuangan :one
INSERT INTO ruangan (
    ruangan_id,
    kantor_id,
    nama_ruangan,
    lantai,
    status,
    created_at,
    created_by
) VALUES (
             @ruangan_id,
             @kantor_id,
             @nama_ruangan,
             @lantai,
             @status::status_enum,
             (now() at time zone 'UTC')::TIMESTAMP,
             @created_by
         )
    RETURNING *;

-- name: GetRuanganByID :one
SELECT *
FROM ruangan
WHERE ruangan_id = @ruangan_id
  AND deleted_at IS NULL;

-- name: ListRuangan :many
SELECT *
FROM ruangan
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateRuangan :one
UPDATE ruangan
SET
    kantor_id = @kantor_id,
    nama_ruangan = @nama_ruangan,
    lantai = @lantai,
    status = @status::status_enum,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE ruangan_id = @ruangan_id
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteRuangan :exec
UPDATE ruangan
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE ruangan_id = @ruangan_id;

-- name: RestoreRuangan :exec
UPDATE ruangan
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE ruangan_id = @ruangan_id;
