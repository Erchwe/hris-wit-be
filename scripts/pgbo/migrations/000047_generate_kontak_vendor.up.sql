DO $$ BEGIN
    CREATE TYPE contact_type_enum AS ENUM ('nomor hp', 'telepon', 'whatsapp', 'email');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE SEQUENCE IF NOT EXISTS kontak_vendor_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS kontak_vendor (
    id BIGINT NOT NULL DEFAULT nextval('kontak_vendor_id_seq'::regclass),
    kontak_vendor_id VARCHAR NOT NULL,
    vendor_id VARCHAR NOT NULL,
    jenis_kontak contact_type_enum,
    isi_kontak VARCHAR(100),
    is_primary BOOLEAN DEFAULT FALSE,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT kontak_vendor_pk PRIMARY KEY (kontak_vendor_id),
    CONSTRAINT kontak_vendor_vendor_fk FOREIGN KEY (vendor_id) REFERENCES vendor(vendor_id)
);