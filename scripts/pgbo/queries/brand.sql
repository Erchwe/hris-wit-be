-- name: CreateBrand :one
INSERT INTO brand (
    brand_id,
    nama_brand,
    status,
    created_at,
    created_by
) VALUES (
             @brand_id,
             @nama_brand,
             @status::status_enum,
             (now() at time zone 'UTC')::TIMESTAMP,
             @created_by
         )
    RETURNING *;

-- name: GetBrandByID :one
SELECT * FROM brand
WHERE brand_id = @brand_id
  AND deleted_at IS NULL;

-- name: ListBrands :many
SELECT * FROM brand
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateBrand :one
UPDATE brand
SET
    nama_brand = @nama_brand,
    status = @status::status_enum,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE brand_id = @brand_id
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteBrand :exec
UPDATE brand
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE brand_id = @brand_id;