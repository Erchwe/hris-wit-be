CREATE SEQUENCE IF NOT EXISTS kategori_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS kategori (
    id BIGINT NOT NULL DEFAULT nextval('kategori_id_seq'),
    kategori_id VARCHAR NOT NULL,
    nama_kategori VARCHAR(100) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT kategori_pk PRIMARY KEY (kategori_id)
);