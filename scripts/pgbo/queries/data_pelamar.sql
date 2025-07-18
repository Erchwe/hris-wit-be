-- name: CreatePelamar :one
INSERT INTO data_pelamar (
    id_data_pelamar, id_lowongan_pekerjaan, email, nama_lengkap, tempat_lahir,
    tgl_lahir, jenis_kelamin, kewarganegaraan, phone, mobile, agama, gol_darah,
    status_menikah, no_ktp, no_npwp, status, asal_kota, gaji_terakhir, harapan_gaji,
    sedang_bekerja, ketersediaan_bekerja, sumber_informasi, alasan, ketersediaan_inter,
    profesi_kerja, created_at, created_by
) VALUES (
    @id_data_pelamar, @id_lowongan_pekerjaan, @email, @nama_lengkap, @tempat_lahir,
    @tgl_lahir, @jenis_kelamin, @kewarganegaraan, @phone, @mobile, @agama, @gol_darah,
    @status_menikah, @no_ktp, @no_npwp, @status, @asal_kota, @gaji_terakhir, @harapan_gaji,
    @sedang_bekerja, @ketersediaan_bekerja, @sumber_informasi, @alasan, @ketersediaan_inter,
    @profesi_kerja, (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetPelamarByID :one
SELECT * FROM data_pelamar WHERE id_data_pelamar = @id_data_pelamar AND deleted_at IS NULL;

-- name: ListPelamar :many
SELECT * FROM data_pelamar WHERE deleted_at IS NULL ORDER BY created_at DESC;

-- name: GetKeluargaByPelamarID :one
SELECT * FROM pelamar_keluarga
WHERE id_pelamar = @id_data_pelamar AND deleted_at IS NULL;

-- name: ListAnakByPelamarID :many
SELECT * FROM pelamar_anak
WHERE id_pelamar = @id_data_pelamar AND deleted_at IS NULL;

-- name: ListSaudaraByPelamarID :many
SELECT * FROM pelamar_saudara_kandung
WHERE id_pelamar = @id_data_pelamar AND deleted_at IS NULL;

-- name: GetPendidikanFormalByPelamarID :many
SELECT * FROM pelamar_pendidikan_formal
WHERE id_pelamar = @id_data_pelamar AND deleted_at IS NULL;

-- name: ListPendidikanNonFormalByPelamarID :many
SELECT * FROM pelamar_pendidikan_non_formal
WHERE id_pelamar = @id_data_pelamar AND deleted_at IS NULL;

-- name: ListBahasaByPelamarID :many
SELECT * FROM pelamar_penguasaan_bahasa
WHERE id_pelamar = @id_data_pelamar AND deleted_at IS NULL;

-- name: GetPengalamanKerjaByPelamarID :many
SELECT * FROM pelamar_pengalaman_kerja
WHERE id_pelamar = @id_data_pelamar AND deleted_at IS NULL;

-- name: ListReferensiByPelamarID :many
SELECT * FROM pelamar_referensi
WHERE id_pelamar = @id_data_pelamar AND deleted_at IS NULL;

-- name: UpdatePelamar :one
UPDATE data_pelamar
SET
    status = @status,
    updated_at = NOW()
WHERE id_data_pelamar = @id_data_pelamar AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeletePelamar :exec
UPDATE data_pelamar
SET deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_data_pelamar = @id_data_pelamar AND deleted_at IS NULL;
