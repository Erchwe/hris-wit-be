-- name: CreatePegawaiPendidikanNonFormal :one
INSERT INTO pegawai_pendidikan_non_formal (
    id_pddk_non_formal, id_pegawai, institusi, jenis_pendidikan, kota,
    created_at, created_by
) VALUES (
    @id_pddk_non_formal, @id_pegawai, @institusi, @jenis_pendidikan, @kota,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPegawaiPendidikanNonFormalByID :one
SELECT * FROM pegawai_pendidikan_non_formal
WHERE id_pddk_non_formal = @id_pddk_non_formal
  AND deleted_at IS NULL;

-- name: ListPegawaiPendidikanNonFormal :many
SELECT * FROM pegawai_pendidikan_non_formal
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdatePegawaiPendidikanNonFormal :one
UPDATE pegawai_pendidikan_non_formal
SET
    institusi = @institusi,
    jenis_pendidikan = @jenis_pendidikan,
    kota = @kota,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_pddk_non_formal = @id_pddk_non_formal
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePegawaiPendidikanNonFormal :exec
UPDATE pegawai_pendidikan_non_formal
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_pddk_non_formal = @id_pddk_non_formal
  AND deleted_at IS NULL;

-- name: RestorePegawaiPendidikanNonFormal :exec
UPDATE pegawai_pendidikan_non_formal
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_pddk_non_formal = @id_pddk_non_formal;
