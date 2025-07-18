-- name: CreatePegawaiAnak :one
INSERT INTO pegawai_anak (
    id_anak, id_pegawai, nama, jenis_kelamin, tempat_lahir,
    tgl_lahir, pendidikan_pekerjaan, created_at, created_by
) VALUES (
    @id_anak, @id_pegawai, @nama, @jenis_kelamin, @tempat_lahir,
    @tgl_lahir, @pendidikan_pekerjaan,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPegawaiAnakByID :one
SELECT * FROM pegawai_anak
WHERE id_anak = @id_anak
  AND deleted_at IS NULL;

-- name: ListPegawaiAnak :many
SELECT * FROM pegawai_anak
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdatePegawaiAnak :one
UPDATE pegawai_anak
SET
    nama = @nama,
    jenis_kelamin = @jenis_kelamin,
    tempat_lahir = @tempat_lahir,
    tgl_lahir = @tgl_lahir,
    pendidikan_pekerjaan = @pendidikan_pekerjaan,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_anak = @id_anak
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePegawaiAnak :exec
UPDATE pegawai_anak
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_anak = @id_anak
  AND deleted_at IS NULL;

-- name: RestorePegawaiAnak :exec
UPDATE pegawai_anak
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_anak = @id_anak;
