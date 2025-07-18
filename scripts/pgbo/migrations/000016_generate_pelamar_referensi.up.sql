CREATE SEQUENCE IF NOT EXISTS pelamar_referensi_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS pelamar_referensi (
                                                 id bigint NOT NULL DEFAULT nextval('pelamar_referensi_id_seq'::regclass),
    id_referensi VARCHAR NOT NULL,
    id_pelamar VARCHAR NOT NULL,
    nama VARCHAR(100) NOT NULL,
    nama_perusahaan VARCHAR(100),
    jabatan VARCHAR(20),
    no_telp_perusahaan VARCHAR(15),
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT pelamar_referensi_pk PRIMARY KEY (id_referensi),
    CONSTRAINT fk_pelamar_referensi FOREIGN KEY (id_pelamar)REFERENCES data_pelamar(id_data_pelamar)
    );