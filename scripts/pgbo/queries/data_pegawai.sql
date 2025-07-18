-- name: CreateDataPegawai :one
INSERT INTO data_pegawai (
    id_data_pegawai, employee_number, divisi, nama_lengkap, tempat_lahir, tgl_lahir,
    jenis_kelamin, kewarganegaraan, phone, mobile, agama, gol_darah, gaji,
    status_menikah, no_ktp, no_npwp, status, created_at, created_by
) VALUES (
    @id_data_pegawai, @employee_number, @divisi, @nama_lengkap, @tempat_lahir, @tgl_lahir,
    @jenis_kelamin, @kewarganegaraan, @phone, @mobile, @agama, @gol_darah, @gaji,
    @status_menikah, @no_ktp, @no_npwp, @status, (now() at time zone 'UTC')::TIMESTAMP, @created_by
) RETURNING *;

-- name: GetDataPegawaiByID :one
SELECT * FROM data_pegawai
WHERE id_data_pegawai = @id_data_pegawai
  AND deleted_at IS NULL;

-- name: ListDataPegawai :many
SELECT * FROM data_pegawai
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateDataPegawai :one
UPDATE data_pegawai
SET
    employee_number = @employee_number,
    divisi = @divisi,
    nama_lengkap = @nama_lengkap,
    tempat_lahir = @tempat_lahir,
    tgl_lahir = @tgl_lahir,
    jenis_kelamin = @jenis_kelamin,
    kewarganegaraan = @kewarganegaraan,
    phone = @phone,
    mobile = @mobile,
    agama = @agama,
    gol_darah = @gol_darah,
    gaji = @gaji,
    status_menikah = @status_menikah,
    no_ktp = @no_ktp,
    no_npwp = @no_npwp,
    status = @status,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = @updated_by
WHERE id_data_pegawai = @id_data_pegawai
  AND deleted_at IS NULL
    RETURNING *;

-- name: SoftDeleteDataPegawai :exec
UPDATE data_pegawai
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = @deleted_by
WHERE id_data_pegawai = @id_data_pegawai
  AND deleted_at IS NULL;

-- name: RestoreDataPegawai :exec
UPDATE data_pegawai
SET
    deleted_at = NULL,
    deleted_by = NULL
WHERE id_data_pegawai = @id_data_pegawai;
