-- name: CreatePelamarPendidikanFormal :one
INSERT INTO pelamar_pendidikan_formal (
    id_pddk_formal, id_pelamar, jenjang_pddk, nama_sekolah, jurusan_fakultas, kota,
    tgl_lulus, ipk, created_at, created_by
) VALUES (
             @id_pddk_formal, @id_pelamar, @jenjang_pddk, @nama_sekolah, @jurusan_fakultas, @kota,
             @tgl_lulus, @ipk, (now() at time zone 'UTC')::TIMESTAMP, @created_by
         ) RETURNING *;

-- name: GetPelamarPendidikanFormalByID :one
SELECT * FROM pelamar_pendidikan_formal
WHERE id_pddk_formal = @id_pddk_formal
  AND deleted_at IS NULL;

-- name: ListPelamarPendidikanFormal :many
SELECT * FROM pelamar_pendidikan_formal
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdatePelamarPendidikanFormal :one
UPDATE pelamar_pendidikan_formal
SET
    jenjang_pddk = @jenjang_pddk,
    nama_sekolah = @nama_sekolah,
    jurusan_fakultas = @jurusan_fakultas,
    kota = @kota,
    tgl_lulus = @tgl_lulus,
    ipk = @ipk,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_pddk_formal = @id_pddk_formal
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePelamarPendidikanFormal :exec
UPDATE pelamar_pendidikan_formal
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_pddk_formal = @id_pddk_formal
  AND deleted_at IS NULL;

-- name: RestorePelamarPendidikanFormal :exec
UPDATE pelamar_pendidikan_formal
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_pddk_formal = @id_pddk_formal;
