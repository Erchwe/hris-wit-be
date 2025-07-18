CREATE SEQUENCE IF NOT EXISTS pegawai_pendidikan_formal_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS pegawai_pendidikan_formal(
                                                        id bigint NOT NULL DEFAULT nextval('pegawai_pendidikan_formal_id_seq'::regclass),
    id_pddk_formal VARCHAR NOT NULL,
    id_pegawai VARCHAR NOT NULL,
    jenjang_pddk VARCHAR(100) NOT NULL,
    nama_sekolah VARCHAR(100) NOT NULL,
    jurusan_fakultas VARCHAR(100),
    kota VARCHAR(20),
    tgl_lulus DATE,
    ipk FLOAT,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT pegawai_pendidikan_formal_pk PRIMARY KEY (id_pddk_formal),
    CONSTRAINT fk_pegawai_pendidikan_formal FOREIGN KEY (id_pegawai) REFERENCES data_pegawai(id_data_pegawai)
    );