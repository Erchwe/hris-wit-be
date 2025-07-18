CREATE SEQUENCE IF NOT EXISTS inventaris_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS inventaris (
    id BIGINT NOT NULL DEFAULT nextval('inventaris_id_seq'::regclass),
    inventaris_id VARCHAR NOT NULL,
    brand_id VARCHAR NOT NULL,
    ruangan_id VARCHAR,
    user_id VARCHAR,
    kategori_id VARCHAR NOT NULL,
    vendor_id VARCHAR NOT NULL,
    nama_inventaris VARCHAR(255) NOT NULL,
    jumlah INT NOT NULL,
    tanggal_beli DATE NOT NULL,
    harga BIGINT NOT NULL,
    keterangan TEXT,
    old_inventory_code VARCHAR(100),
    image_url VARCHAR(255),
    status status_enum,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT inventaris_pk PRIMARY KEY (inventaris_id),
    CONSTRAINT brand_inventaris_fk FOREIGN KEY (brand_id) REFERENCES brand(brand_id),
    CONSTRAINT ruangan_inventaris_fk FOREIGN KEY (ruangan_id) REFERENCES ruangan(ruangan_id),
    CONSTRAINT user_inventaris_fk FOREIGN KEY (user_id) REFERENCES data_pegawai(id_data_pegawai),
    CONSTRAINT kategori_inventaris_fk FOREIGN KEY (kategori_id) REFERENCES kategori(kategori_id)
);