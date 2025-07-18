CREATE SEQUENCE IF NOT EXISTS pegawai_anak_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS pegawai_anak (
                                            id bigint NOT NULL DEFAULT nextval('pegawai_anak_id_seq'::regclass),
    id_anak VARCHAR NOT NULL,
    id_pegawai VARCHAR NOT NULL,
    nama VARCHAR(256) NOT NULL,
    jenis_kelamin jenis_kelamin_enum NOT NULL,
    tempat_lahir VARCHAR(20),
    tgl_lahir DATE,
    pendidikan_pekerjaan VARCHAR(20),
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT pegawai_anak_pk PRIMARY KEY (id_anak),
    CONSTRAINT fk_pegawai_anak FOREIGN KEY (id_pegawai) REFERENCES data_pegawai(id_data_pegawai)
    );
