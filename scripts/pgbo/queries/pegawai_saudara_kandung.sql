-- name: CreatePegawaiSaudaraKandung :one
INSERT INTO pegawai_saudara_kandung (
    id_saudara, id_pegawai, nama, jenis_kelamin, tempat_lahir,
    tgl_lahir, pendidikan_pekerjaan, created_at, created_by
) VALUES (
    @id_saudara, @id_pegawai, @nama, @jenis_kelamin, @tempat_lahir,
    @tgl_lahir, @pendidikan_pekerjaan,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPegawaiSaudaraKandungByID :one
SELECT * FROM pegawai_saudara_kandung
WHERE id_saudara = @id_saudara
  AND deleted_at IS NULL;

-- name: ListPegawaiSaudaraKandung :many
SELECT * FROM pegawai_saudara_kandung
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdatePegawaiSaudaraKandung :one
UPDATE pegawai_saudara_kandung
SET
    nama = @nama,
    jenis_kelamin = @jenis_kelamin,
    tempat_lahir = @tempat_lahir,
    tgl_lahir = @tgl_lahir,
    pendidikan_pekerjaan = @pendidikan_pekerjaan,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_saudara = @id_saudara
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePegawaiSaudaraKandung :exec
UPDATE pegawai_saudara_kandung
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_saudara = @id_saudara
  AND deleted_at IS NULL;

-- name: RestorePegawaiSaudaraKandung :exec
UPDATE pegawai_saudara_kandung
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_saudara = @id_saudara;
