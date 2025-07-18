-- name: CreatePeminjaman :one
INSERT INTO peminjaman (
    peminjaman_id,
    inventaris_id,
    user_id,
    tgl_pinjam,
    tgl_kembali,
    status_peminjaman,
    notes,
    created_at,
    created_by
) VALUES (
             @peminjaman_id,
             @inventaris_id,
             @user_id,
             @tgl_pinjam,
             @tgl_kembali,
             @status_peminjaman::status_peminjaman_enum,
             @notes,
             (now() at time zone 'UTC')::TIMESTAMP,
             @created_by
         )
    RETURNING *;

-- name: GetPeminjamanByID :one
SELECT
    p.peminjaman_id,
    p.tgl_pinjam,
    p.tgl_kembali,
    p.status_peminjaman,
    p.notes,
    i.nama_inventaris,
    ub.name AS nama_peminjam_user
FROM peminjaman p
         LEFT JOIN inventaris i ON p.inventaris_id = i.inventaris_id
         LEFT JOIN user_backoffice ub ON p.user_id = ub.guid
WHERE p.peminjaman_id = @peminjaman_id AND p.deleted_at IS NULL;

-- name: ListPeminjaman :many
SELECT
    p.peminjaman_id,
    p.tgl_pinjam,
    p.tgl_kembali,
    p.status_peminjaman,
    p.notes,
    i.nama_inventaris,
    ub.name AS nama_peminjam_user
FROM peminjaman p
         LEFT JOIN inventaris i ON p.inventaris_id = i.inventaris_id
         LEFT JOIN user_backoffice ub ON p.user_id = ub.guid
WHERE p.deleted_at IS NULL
ORDER BY p.created_at DESC;

-- name: UpdatePeminjaman :one
UPDATE peminjaman
SET
    tgl_pinjam = @tgl_pinjam,               -- Kembali ke nama kolom
    tgl_kembali = @tgl_kembali,             -- Kembali ke nama kolom
    status_peminjaman = @status_peminjaman::status_peminjaman_enum, -- Kembali ke nama kolom
    notes = @notes,                         -- Kembali ke nama kolom
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by                -- Kembali ke nama kolom
WHERE peminjaman_id = @peminjaman_id AND deleted_at IS NULL -- Kembali ke nama kolom
    RETURNING *;

-- name: UpdatePeminjamanStatus :one
UPDATE peminjaman
SET
    status_peminjaman = @status_peminjaman::status_peminjaman_enum, -- Kembali ke nama kolom
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by -- Kembali ke nama kolom
WHERE peminjaman_id = @peminjaman_id AND deleted_at IS NULL -- Kembali ke nama kolom
    RETURNING *;

-- name: SoftDeletePeminjaman :exec
UPDATE peminjaman
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE peminjaman_id = @peminjaman_id;

-- name: RestorePeminjaman :exec
UPDATE peminjaman
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE peminjaman_id = @peminjaman_id;

-- name: ListPeminjamanByUserID :many
SELECT
    p.peminjaman_id,
    p.tgl_pinjam,
    p.tgl_kembali,
    p.status_peminjaman,
    p.notes,
    i.nama_inventaris -- Ambil nama inventaris
FROM peminjaman p
         JOIN inventaris i ON p.inventaris_id = i.inventaris_id
WHERE p.user_id = @user_id AND p.deleted_at IS NULL
ORDER BY p.created_at DESC;

-- name: ListPendingPeminjaman :many
SELECT
    p.peminjaman_id,
    p.tgl_pinjam,
    p.tgl_kembali,
    p.notes,
    i.nama_inventaris,
    ub.name AS nama_peminjam_user, -- Nama user yang mengajukan
    ub.guid AS user_id_peminjam -- ID user yang mengajukan
FROM peminjaman p
         JOIN inventaris i ON p.inventaris_id = i.inventaris_id
         JOIN user_backoffice ub ON p.user_id = ub.guid
WHERE p.status_peminjaman = 'Menunggu Persetujuan' AND p.deleted_at IS NULL
ORDER BY p.created_at ASC;

-- name: ListOverduePeminjamanByUserID :many
SELECT
    p.peminjaman_id,
    i.nama_inventaris,
    p.tgl_kembali, -- Ini adalah tanggal kembali rencana
    p.tgl_pinjam,
    p.status_peminjaman
FROM peminjaman p
         JOIN inventaris i ON p.inventaris_id = i.inventaris_id
WHERE p.user_id = @user_id
  AND p.status_peminjaman = 'Sedang Dipinjam'
  AND p.tgl_kembali < (now() at time zone 'UTC')::DATE -- Tanggal kembali sudah lewat
  AND p.deleted_at IS NULL
ORDER BY p.tgl_kembali ASC;

-- name: ListAvailableInventaris :many
SELECT
    i.inventaris_id,
    i.nama_inventaris,
    i.status, -- Status inventaris (Aktif/Tidak Aktif)
    i.keterangan
FROM inventaris i
         LEFT JOIN peminjaman p ON i.inventaris_id = p.inventaris_id AND p.deleted_at IS NULL AND p.status_peminjaman = 'Sedang Dipinjam'
WHERE i.deleted_at IS NULL
  AND i.status = 'Aktif'
  AND p.peminjaman_id IS NULL
ORDER BY i.nama_inventaris ASC;