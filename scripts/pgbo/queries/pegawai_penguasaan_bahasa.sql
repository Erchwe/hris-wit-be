-- name: CreatePegawaiPenguasaanBahasa :one
INSERT INTO pegawai_penguasaan_bahasa (
    id_bahasa, id_pegawai, bahasa, lisan, tulisan, keterangan,
    created_at, created_by
) VALUES (
    @id_bahasa, @id_pegawai, @bahasa, @lisan, @tulisan, @keterangan,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPegawaiPenguasaanBahasaByID :one
SELECT * FROM pegawai_penguasaan_bahasa
WHERE id_bahasa = @id_bahasa
  AND deleted_at IS NULL;

-- name: ListPegawaiPenguasaanBahasa :many
SELECT * FROM pegawai_penguasaan_bahasa
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdatePegawaiPenguasaanBahasa :one
UPDATE pegawai_penguasaan_bahasa
SET
    bahasa = @bahasa,
    lisan = @lisan,
    tulisan = @tulisan,
    keterangan = @keterangan,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_bahasa = @id_bahasa
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePegawaiPenguasaanBahasa :exec
UPDATE pegawai_penguasaan_bahasa
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_bahasa = @id_bahasa
  AND deleted_at IS NULL;

-- name: RestorePegawaiPenguasaanBahasa :exec
UPDATE pegawai_penguasaan_bahasa
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_bahasa = @id_bahasa;
