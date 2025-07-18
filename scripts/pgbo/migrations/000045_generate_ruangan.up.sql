CREATE SEQUENCE IF NOT EXISTS ruangan_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS ruangan (
    id BIGINT NOT NULL DEFAULT nextval('ruangan_id_seq'::regclass),
    ruangan_id VARCHAR NOT NULL,
    kantor_id VARCHAR NOT NULL,
    nama_ruangan VARCHAR(100) NOT NULL,
    lantai INT,
    status status_enum,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT ruangan_pk PRIMARY KEY (ruangan_id),
    CONSTRAINT ruangan_kantor_fk FOREIGN KEY (kantor_id) REFERENCES kantor(kantor_id)
);