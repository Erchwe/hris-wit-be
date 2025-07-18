-- name: CreatePelamarPengalamanKerja :one
INSERT INTO pelamar_pengalaman_kerja (
    id_pengalaman_kerja, id_pelamar, nama_perusahaan, periode, jabatan, gaji, alasan_pindah,
    created_at, created_by
) VALUES (
    @id_pengalaman_kerja, @id_pelamar, @nama_perusahaan, @periode, @jabatan, @gaji, @alasan_pindah,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPelamarPengalamanKerjaByID :one
SELECT * FROM pelamar_pengalaman_kerja
WHERE id_pengalaman_kerja = @id_pengalaman_kerja
  AND deleted_at IS NULL;

-- name: ListPelamarPengalamanKerja :many
SELECT * FROM pelamar_pengalaman_kerja
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdatePelamarPengalamanKerja :one
UPDATE pelamar_pengalaman_kerja
SET
    nama_perusahaan = @nama_perusahaan,
    periode = @periode,
    jabatan = @jabatan,
    gaji = @gaji,
    alasan_pindah = @alasan_pindah,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_pengalaman_kerja = @id_pengalaman_kerja
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePelamarPengalamanKerja :exec
UPDATE pelamar_pengalaman_kerja
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_pengalaman_kerja = @id_pengalaman_kerja
  AND deleted_at IS NULL;

-- name: RestorePelamarPengalamanKerja :exec
UPDATE pelamar_pengalaman_kerja
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_pengalaman_kerja = @id_pengalaman_kerja;
