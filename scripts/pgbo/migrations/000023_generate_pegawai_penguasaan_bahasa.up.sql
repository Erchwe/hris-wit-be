CREATE SEQUENCE IF NOT EXISTS pegawai_penguasaan_bahasa_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS pegawai_penguasaan_bahasa(
                                                        id bigint NOT NULL DEFAULT nextval('pegawai_penguasaan_bahasa_id_seq'::regclass),
    id_bahasa VARCHAR NOT NULL,
    id_pegawai VARCHAR NOT NULL,
    bahasa VARCHAR(50) NOT NULL,
    lisan VARCHAR(20),
    tulisan VARCHAR(20),
    keterangan TEXT,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT pegawai_penguasaan_bahasa_pk PRIMARY KEY (id_bahasa),
    CONSTRAINT fk_pegawai_penguasaan_bahasa FOREIGN KEY (id_pegawai) REFERENCES data_pegawai(id_data_pegawai)
    );