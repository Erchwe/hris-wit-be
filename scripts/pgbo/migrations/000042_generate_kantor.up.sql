DO $$ BEGIN
    CREATE TYPE status_enum AS ENUM ('Aktif', 'Tidak Aktif');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE SEQUENCE IF NOT EXISTS kantor_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS kantor (
    id BIGINT NOT NULL DEFAULT nextval('kantor_id_seq'::regclass),
    kantor_id VARCHAR NOT NULL,
    nama_kantor VARCHAR(255) NOT NULL,
    kota VARCHAR(100) NOT NULL,
    alamat TEXT NOT NULL,
    nomor_telp VARCHAR(20) NOT NULL,
    status status_enum,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT kantor_pk PRIMARY KEY (kantor_id)
);