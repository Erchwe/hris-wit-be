-- name: CreateInventaris :one
INSERT INTO inventaris (
    inventaris_id,
    brand_id,
    ruangan_id,
    user_id,
    kategori_id,
    vendor_id,
    nama_inventaris,
    jumlah,
    tanggal_beli,
    harga,
    keterangan,
    old_inventory_code,
    image_url,
    status,
    created_at,
    created_by
) VALUES (
    @inventaris_id,
    @brand_id,
    @ruangan_id,
    @user_id,
    @kategori_id,
    @vendor_id,
    @nama_inventaris,
    @jumlah,
    @tanggal_beli,
    @harga,
    @keterangan,
    @old_inventory_code,
    @image_url,
    @status::status_enum,
    (now() at time zone 'UTC')::TIMESTAMP,
    @created_by
)
RETURNING *;

-- name: GetInventarisByID :one
SELECT * FROM inventaris
WHERE inventaris_id = @inventaris_id AND deleted_at IS NULL;

-- name: ListInventaris :many
SELECT * FROM inventaris
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateInventaris :one
UPDATE inventaris
SET
    brand_id = @brand_id,
    ruangan_id = @ruangan_id,
    user_id = @user_id,
    kategori_id = @kategori_id,
    vendor_id = @vendor_id,
    nama_inventaris = @nama_inventaris,
    jumlah = @jumlah,
    tanggal_beli = @tanggal_beli,
    harga = @harga,
    keterangan = @keterangan,
    old_inventory_code = @old_inventory_code,
    image_url = @image_url,
    status = @status::status_enum,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE inventaris_id = @inventaris_id AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteInventaris :exec
UPDATE inventaris
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE inventaris_id = @inventaris_id;

-- name: RestoreInventaris :exec
UPDATE inventaris
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE inventaris_id = @inventaris_id;

-- name: ListInventarisWithRelations :many
SELECT
    i.inventaris_id,
    i.nama_inventaris,
    i.tanggal_beli,
    i.harga,
    i.image_url,
    i.jumlah,
    i.keterangan,
    i.old_inventory_code,
    i.status,

    -- relasi vendor
    v.vendor_id,
    v.nama_vendor,

    -- relasi brand
    b.brand_id,
    b.nama_brand,

    -- relasi kategori
    k.kategori_id,
    k.nama_kategori,

    -- relasi ruangan
    r.ruangan_id,
    r.nama_ruangan

FROM inventaris i
         LEFT JOIN vendor v ON i.vendor_id = v.vendor_id
         LEFT JOIN brand b ON i.brand_id = b.brand_id
         LEFT JOIN kategori k ON i.kategori_id = k.kategori_id
         LEFT JOIN ruangan r ON i.ruangan_id = r.ruangan_id
WHERE i.deleted_at IS NULL
ORDER BY i.created_at DESC;

-- name: GetInventarisWithRelationsByID :one
SELECT
    i.inventaris_id,
    i.nama_inventaris,
    i.tanggal_beli,
    i.harga,
    i.image_url,
    i.jumlah,
    i.keterangan,
    i.old_inventory_code,
    i.status,

    i.vendor_id,
    v.nama_vendor,

    i.brand_id,
    b.nama_brand,

    i.kategori_id,
    k.nama_kategori,

    i.ruangan_id,
    r.nama_ruangan

FROM inventaris i
         LEFT JOIN vendor v ON v.vendor_id = i.vendor_id
         LEFT JOIN brand b ON b.brand_id = i.brand_id
         LEFT JOIN kategori k ON k.kategori_id = i.kategori_id
         LEFT JOIN ruangan r ON r.ruangan_id = i.ruangan_id

WHERE i.inventaris_id = $1
  AND i.deleted_at IS NULL;
