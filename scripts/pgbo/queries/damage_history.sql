-- name: CreateDamageHistory :one
INSERT INTO damage_history (
    damage_history_id,
    inventaris_id,
    id_pegawai,
    tgl_rusak,
    tgl_awal_perbaikan,
    tgl_selesai_perbaikan,
    description,
    biaya_perbaikan,
    vendor_perbaikan,
    status,
    created_at,
    created_by
) VALUES (
             @damage_history_id,
             @inventaris_id,
             @id_pegawai,
             @tgl_rusak,
             @tgl_awal_perbaikan,
             @tgl_selesai_perbaikan,
             @description,
             @biaya_perbaikan,
             @vendor_perbaikan,
             @status::status_enum,
             (now() at time zone 'UTC')::TIMESTAMP,
             @created_by
         )
    RETURNING *;

-- name: GetDamageHistoryByID :one
SELECT *
FROM damage_history
WHERE damage_history_id = @damage_history_id
  AND deleted_at IS NULL;

-- name: ListDamageHistory :many
SELECT *
FROM damage_history
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateDamageHistory :one
UPDATE damage_history
SET
    inventaris_id = @inventaris_id,
    id_pegawai = @id_pegawai,
    tgl_rusak = @tgl_rusak,
    tgl_awal_perbaikan = @tgl_awal_perbaikan,
    tgl_selesai_perbaikan = @tgl_selesai_perbaikan,
    description = @description,
    biaya_perbaikan = @biaya_perbaikan,
    vendor_perbaikan = @vendor_perbaikan,
    status = @status::status_enum,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE damage_history_id = @damage_history_id
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteDamageHistory :exec
UPDATE damage_history
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE damage_history_id = @damage_history_id;

-- name: RestoreDamageHistory :exec
UPDATE damage_history
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE damage_history_id = @damage_history_id;
