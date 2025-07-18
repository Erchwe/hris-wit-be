CREATE SEQUENCE IF NOT EXISTS brand_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS brand (
    id BIGINT NOT NULL DEFAULT nextval('brand_id_seq'::regclass),
    brand_id VARCHAR NOT NULL,
    nama_brand VARCHAR(100) NOT NULL,
    status status_enum,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT brand_pk PRIMARY KEY (brand_id)
);