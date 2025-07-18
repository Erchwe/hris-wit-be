-- name: CreatePelamarKeluarga :one
INSERT INTO pelamar_keluarga (
    id_keluarga, id_pelamar, nama_istri_suami, jenis_kelamin, tempat_lahir,
    tgl_lahir, pendidikan_terakhir, pekerjaan_skrg, alamat_rumah,
    created_at, created_by
) VALUES (
    @id_keluarga, @id_pelamar, @nama_istri_suami, @jenis_kelamin, @tempat_lahir,
    @tgl_lahir, @pendidikan_terakhir, @pekerjaan_skrg, @alamat_rumah,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPelamarKeluargaByID :one
SELECT * FROM pelamar_keluarga WHERE id_keluarga = @id_keluarga AND deleted_at IS NULL;

-- name: ListPelamarKeluarga :many
SELECT * FROM pelamar_keluarga WHERE deleted_at IS NULL ORDER BY created_at DESC;

-- name: UpdatePelamarKeluarga :one
UPDATE pelamar_keluarga
SET
    nama_istri_suami = @nama_istri_suami,
    jenis_kelamin = @jenis_kelamin,
    tempat_lahir = @tempat_lahir,
    tgl_lahir = @tgl_lahir,
    pendidikan_terakhir = @pendidikan_terakhir,
    pekerjaan_skrg = @pekerjaan_skrg,
    alamat_rumah = @alamat_rumah,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_keluarga = @id_keluarga AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePelamarKeluarga :exec
UPDATE pelamar_keluarga
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_keluarga = @id_keluarga AND deleted_at IS NULL;

-- name: RestorePelamarKeluarga :exec
UPDATE pelamar_keluarga
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_keluarga = @id_keluarga;
