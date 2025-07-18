-- name: CreatePelamarPendidikanNonFormal :one
INSERT INTO pelamar_pendidikan_non_formal (
    id_pddk_non_formal, id_pelamar, institusi, jenis_pendidikan, kota,
    tgl_lulus, created_at, created_by
) VALUES (
    @id_pddk_non_formal, @id_pelamar, @institusi, @jenis_pendidikan, @kota,
    @tgl_lulus, (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPelamarPendidikanNonFormalByID :one
SELECT * FROM pelamar_pendidikan_non_formal
WHERE id_pddk_non_formal = @id_pddk_non_formal
  AND deleted_at IS NULL;

-- name: ListPelamarPendidikanNonFormal :many
SELECT * FROM pelamar_pendidikan_non_formal
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdatePelamarPendidikanNonFormal :one
UPDATE pelamar_pendidikan_non_formal
SET
    institusi = @institusi,
    jenis_pendidikan = @jenis_pendidikan,
    kota = @kota,
    tgl_lulus = @tgl_lulus,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_pddk_non_formal = @id_pddk_non_formal
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePelamarPendidikanNonFormal :exec
UPDATE pelamar_pendidikan_non_formal
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_pddk_non_formal = @id_pddk_non_formal
  AND deleted_at IS NULL;

-- name: RestorePelamarPendidikanNonFormal :exec
UPDATE pelamar_pendidikan_non_formal
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_pddk_non_formal = @id_pddk_non_formal;
