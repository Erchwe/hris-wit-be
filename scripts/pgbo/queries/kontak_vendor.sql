-- name: CreateKontakVendor :one
INSERT INTO kontak_vendor (
    kontak_vendor_id,
    vendor_id,
    jenis_kontak,
    isi_kontak,
    is_primary,
    created_at,
    created_by
) VALUES (
             @kontak_vendor_id,
             @vendor_id,
             @jenis_kontak::contact_type_enum,
             @isi_kontak,
             @is_primary,
             (now() at time zone 'UTC')::TIMESTAMP,
             @created_by
         )
    RETURNING *;

-- name: GetKontakVendorByID :one
SELECT *
FROM kontak_vendor
WHERE kontak_vendor_id = @kontak_vendor_id
  AND deleted_at IS NULL;

-- name: ListKontakVendorByVendorID :many
SELECT *
FROM kontak_vendor
WHERE vendor_id = @vendor_id
  AND deleted_at IS NULL
ORDER BY is_primary DESC, created_at DESC;

-- name: UpdateKontakVendor :one
UPDATE kontak_vendor
SET
    jenis_kontak = @jenis_kontak::contact_type_enum,
    isi_kontak = @isi_kontak,
    is_primary = @is_primary,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE kontak_vendor_id = @kontak_vendor_id
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteKontakVendor :exec
UPDATE kontak_vendor
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE kontak_vendor_id = @kontak_vendor_id;

-- name: SoftDeleteKontakVendorByVendorID :exec
UPDATE kontak_vendor
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE vendor_id = @vendor_id
  AND deleted_at IS NULL;

-- name: RestoreKontakVendor :exec
UPDATE kontak_vendor
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE kontak_vendor_id = @kontak_vendor_id;
