CREATE SEQUENCE IF NOT EXISTS lowongan_pekerjaan_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS lowongan_pekerjaan
(
    id bigint NOT NULL DEFAULT nextval('lowongan_pekerjaan_id_seq'::regclass),
    id_lowongan_pekerjaan VARCHAR NOT NULL,
    posisi VARCHAR(100) NOT NULL,
    tgl_buka_lowongan DATE NOT NULL,
    tgl_tutup_lowongan DATE NOT NULL,
    kriteria TEXT,
    deskripsi TEXT,
    link_lowongan VARCHAR(256),
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT lowongan_pekerjaan_pk PRIMARY KEY (id_lowongan_pekerjaan)
);

INSERT INTO lowongan_pekerjaan (
    id_lowongan_pekerjaan,
    posisi,
    tgl_buka_lowongan,
    tgl_tutup_lowongan,
    kriteria,
    deskripsi,
    link_lowongan,
    created_at,
    created_by
) VALUES (
'LOW001',
'Backend Developer',
'2025-05-01',
'2025-06-01',
'Pengalaman minimal 1 tahun',
'Bekerja dengan tim, menggunakan Golang dan PostgreSQL',
'https://example.com/job/backend-dev',
NOW(),
'seeder'
);

