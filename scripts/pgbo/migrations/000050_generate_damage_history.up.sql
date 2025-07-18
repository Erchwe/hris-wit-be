CREATE SEQUENCE IF NOT EXISTS damage_history_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS damage_history (
    id BIGINT NOT NULL DEFAULT nextval('damage_history_id_seq'::regclass),
    damage_history_id VARCHAR NOT NULL,
    inventaris_id VARCHAR NOT NULL,
    id_pegawai VARCHAR NOT NULL,
    tgl_rusak DATE NOT NULL,
    tgl_awal_perbaikan DATE NOT NULL,
    tgl_selesai_perbaikan DATE NOT NULL,
    description TEXT,
    biaya_perbaikan BIGINT NOT NULL,
    vendor_perbaikan BIGINT NOT NULL,
    status status_enum,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT damage_history_pk PRIMARY KEY (damage_history_id),
    CONSTRAINT inventaris_damage_history_fk FOREIGN KEY (inventaris_id) REFERENCES inventaris(inventaris_id),
    CONSTRAINT user_damage_history_fk FOREIGN KEY (id_pegawai) REFERENCES data_pegawai(id_data_pegawai)
);