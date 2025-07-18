-- name: CreateVendor :one
INSERT INTO vendor (
    vendor_id,
    nama_vendor,
    alamat,
    status,
    created_at,
    created_by
) VALUES (
             @vendor_id,
             @nama_vendor,
             @alamat,
             @status::status_enum,
             (now() at time zone 'UTC')::TIMESTAMP,
             @created_by
         )
    RETURNING *;

-- name: GetVendorByID :one
SELECT *
FROM vendor
WHERE vendor_id = @vendor_id
  AND deleted_at IS NULL;

-- name: ListVendors :many
SELECT *
FROM vendor
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateVendor :one
UPDATE vendor
SET
    nama_vendor = @nama_vendor,
    alamat = @alamat,
    status = @status::status_enum,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE vendor_id = @vendor_id
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteVendor :exec
UPDATE vendor
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE vendor_id = @vendor_id;

-- name: RestoreVendor :exec
UPDATE vendor
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE vendor_id = @vendor_id;
