DO $$ BEGIN
CREATE TYPE jenis_kelamin_enum AS ENUM ('Laki-laki', 'Perempuan');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
CREATE TYPE status_enum AS ENUM (
    'New',
    'Short list',
    'HR Interview',
    'User Interview',
    'Refference Checking',
    'Offering',
    'Psikotest',
    'Hired',
    'Rejected'
);
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
CREATE TYPE profesi_enum AS ENUM ('Junior', 'Senior', 'Manager');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE SEQUENCE IF NOT EXISTS data_pelamar_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS data_pelamar(
    id bigint NOT NULL DEFAULT nextval('data_pelamar_id_seq'::regclass),
    id_data_pelamar VARCHAR NOT NULL,
    id_lowongan_pekerjaan VARCHAR NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    nama_lengkap VARCHAR(256) NOT NULL,
    tempat_lahir VARCHAR(20),
    tgl_lahir DATE,
    jenis_kelamin jenis_kelamin_enum NOT NULL,
    kewarganegaraan VARCHAR(20),
    phone VARCHAR(15),
    mobile VARCHAR(15),
    agama VARCHAR(10),
    gol_darah VARCHAR(2),
    status_menikah BOOLEAN,
    no_ktp VARCHAR(16) UNIQUE,
    no_npwp VARCHAR(16),
    status status_enum,
    asal_kota VARCHAR(20),
    gaji_terakhir MONEY,
    harapan_gaji MONEY,
    sedang_bekerja VARCHAR(20),
    ketersediaan_bekerja DATE,
    sumber_informasi VARCHAR(20),
    alasan TEXT,
    ketersediaan_inter TIMESTAMP,
    profesi_kerja profesi_enum,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT data_pelamar_pk PRIMARY KEY (id_data_pelamar),
    CONSTRAINT fk_data_pelamar FOREIGN KEY (id_lowongan_pekerjaan)REFERENCES lowongan_pekerjaan(id_lowongan_pekerjaan)
    );

CREATE UNIQUE INDEX nama_idx
    ON data_pelamar (email);