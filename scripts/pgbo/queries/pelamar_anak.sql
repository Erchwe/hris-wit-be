-- name: CreatePelamarAnak :one
INSERT INTO pelamar_anak (
    id_anak, id_pelamar, nama, jenis_kelamin, tempat_lahir,
    pendidikan_pekerjaan, tgl_lahir, created_at, created_by
) VALUES (
    @id_anak, @id_pelamar, @nama, @jenis_kelamin, @tempat_lahir,
    @pendidikan_pekerjaan, @tgl_lahir, (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPelamarAnakByID :one
SELECT * FROM pelamar_anak WHERE id_anak = @id_anak AND deleted_at IS NULL;

-- name: ListPelamarAnak :many
SELECT * FROM pelamar_anak WHERE deleted_at IS NULL ORDER BY created_at DESC;

-- name: UpdatePelamarAnak :one
UPDATE pelamar_anak
SET
    nama = @nama,
    jenis_kelamin = @jenis_kelamin,
    tempat_lahir = @tempat_lahir,
    pendidikan_pekerjaan = @pendidikan_pekerjaan,
    tgl_lahir = @tgl_lahir,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_anak = @id_anak AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePelamarAnak :exec
UPDATE pelamar_anak
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_anak = @id_anak AND deleted_at IS NULL;

-- name: RestorePelamarAnak :exec
UPDATE pelamar_anak
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_anak = @id_anak;