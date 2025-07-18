CREATE SEQUENCE IF NOT EXISTS client_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS client (
    id BIGINT NOT NULL DEFAULT nextval('client_id_seq'::regclass),
    client_id VARCHAR NOT NULL,
    client_name VARCHAR(255),
    shipment_address VARCHAR(255),
    billing_address VARCHAR(255),
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT client_pk PRIMARY KEY (client_id)
);
