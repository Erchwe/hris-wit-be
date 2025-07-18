CREATE SEQUENCE IF NOT EXISTS vendor_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS vendor (
    id BIGINT NOT NULL DEFAULT nextval('vendor_id_seq'),
    vendor_id VARCHAR NOT NULL,
    nama_vendor VARCHAR(100) NOT NULL,
    alamat TEXT NOT NULL,
    status status_enum,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT vendor_pk PRIMARY KEY (vendor_id)
);