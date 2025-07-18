-- name: GetInventarisStats :one
SELECT
    COUNT(i.inventaris_id) AS total_inventaris,
    COUNT(CASE
              WHEN i.status = 'Tidak Aktif' THEN 1
              ELSE NULL
        END) AS rusak_maintenance,
    COUNT(CASE
              WHEN i.status = 'Aktif' AND p.status_peminjaman = 'Sedang Dipinjam' THEN 1
              ELSE NULL
        END) AS sedang_dipinjam,
    COUNT(CASE
              WHEN i.status = 'Aktif' AND p.status_peminjaman IS NULL THEN 1
              ELSE NULL
        END) AS tersedia
FROM inventaris i
         LEFT JOIN peminjaman p ON i.inventaris_id = p.inventaris_id AND p.deleted_at IS NULL
WHERE i.deleted_at IS NULL;

-- name: ListRecentActivities :many
SELECT
    p.peminjaman_id,
    p.tgl_pinjam,
    p.tgl_kembali,
    p.status_peminjaman, -- Ini adalah nilai ENUM aslinya
    i.nama_inventaris,
    ub.name AS nama_peminjam_user,
    CASE
        -- Jika 'Tidak Dipinjam' berarti sudah selesai/dikembalikan:
        WHEN p.status_peminjaman = 'Tidak Dipinjam' THEN 'Dikembalikan'
        -- Jika 'Sedang Dipinjam' berarti masih dipinjam:
        WHEN p.status_peminjaman = 'Sedang Dipinjam' THEN 'Dipinjam'
        -- Tambahkan kasus untuk 'Menunggu Persetujuan' jika perlu ditampilkan di status_display
        WHEN p.status_peminjaman = 'Menunggu Persetujuan' THEN 'Menunggu Persetujuan' -- Menggunakan string yang cocok
        ELSE 'Lainnya' -- Fallback jika ada status lain yang tidak ditangani
        END AS status_display
FROM peminjaman p
         LEFT JOIN inventaris i ON p.inventaris_id = i.inventaris_id
         LEFT JOIN user_backoffice ub ON p.user_id = ub.guid
WHERE p.deleted_at IS NULL
ORDER BY p.created_at DESC
    LIMIT 5;

-- name: ListRecentPeminjam :many
SELECT
    ub.guid AS user_id,
    ub.name AS nama_peminjam,
    MAX(p.tgl_pinjam) AS tanggal_terakhir_pinjam
FROM peminjaman p
         JOIN user_backoffice ub ON p.user_id = ub.guid
WHERE p.deleted_at IS NULL
GROUP BY ub.guid, ub.name
ORDER BY tanggal_terakhir_pinjam DESC
    LIMIT 2;

-- name: ListNotReturnedInventaris :many
SELECT
    p.peminjaman_id,
    i.nama_inventaris,
    p.tgl_pinjam,
    p.tgl_kembali AS tanggal_kembali_rencana,
    ub.name AS nama_peminjam
FROM peminjaman p
         JOIN inventaris i ON p.inventaris_id = i.inventaris_id
         JOIN user_backoffice ub ON p.user_id = ub.guid
WHERE p.deleted_at IS NULL
  AND p.status_peminjaman = 'Sedang Dipinjam' -- Hanya yang statusnya masih 'Sedang Dipinjam'
  AND p.tgl_kembali < (now() at time zone 'UTC')::DATE -- Dan tanggal kembali rencana sudah lewat
ORDER BY p.tgl_kembali ASC;

-- name: ListNewVendors :many
SELECT
    v.vendor_id,
    v.nama_vendor,
    kv.isi_kontak AS kontak_vendor,
    kv.jenis_kontak,
    v.created_at
FROM vendor v
         LEFT JOIN kontak_vendor kv ON v.vendor_id = kv.vendor_id AND kv.is_primary = TRUE AND kv.deleted_at IS NULL
WHERE v.deleted_at IS NULL
ORDER BY v.created_at DESC
    LIMIT 2;