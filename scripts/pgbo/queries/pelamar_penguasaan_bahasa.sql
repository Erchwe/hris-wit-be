-- name: CreatePelamarPenguasaanBahasa :one
INSERT INTO pelamar_penguasaan_bahasa (
    id_bahasa, id_pelamar, bahasa, lisan, tulisan, keterangan,
    created_at, created_by
) VALUES (
    @id_bahasa, @id_pelamar, @bahasa, @lisan, @tulisan, @keterangan,
    (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPelamarPenguasaanBahasaByID :one
SELECT * FROM pelamar_penguasaan_bahasa
WHERE id_bahasa = @id_bahasa
  AND deleted_at IS NULL;

-- name: ListPelamarPenguasaanBahasa :many
SELECT * FROM pelamar_penguasaan_bahasa
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdatePelamarPenguasaanBahasa :one
UPDATE pelamar_penguasaan_bahasa
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

-- name: SoftDeletePelamarPenguasaanBahasa :exec
UPDATE pelamar_penguasaan_bahasa
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_bahasa = @id_bahasa
  AND deleted_at IS NULL;

-- name: RestorePelamarPenguasaanBahasa :exec
UPDATE pelamar_penguasaan_bahasa
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_bahasa = @id_bahasa;
