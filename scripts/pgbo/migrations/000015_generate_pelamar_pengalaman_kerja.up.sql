CREATE SEQUENCE IF NOT EXISTS pelamar_pengalaman_kerja_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS pelamar_pengalaman_kerja (
    id bigint NOT NULL DEFAULT nextval('pelamar_pengalaman_kerja_id_seq'::regclass),
    id_pengalaman_kerja VARCHAR NOT NULL,
    id_pelamar VARCHAR NOT NULL,
    nama_perusahaan VARCHAR(100) NOT NULL,
    periode VARCHAR(5),
    jabatan VARCHAR(20),
    gaji MONEY,
    alasan_pindah VARCHAR(30),
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT pelamar_pengalaman_kerja_pk PRIMARY KEY (id_pengalaman_kerja),
    CONSTRAINT fk_pelamar_pengalama_kerja FOREIGN KEY (id_pelamar)REFERENCES data_pelamar(id_data_pelamar)
    );