DO $$ BEGIN
CREATE TYPE jenis_kelamin_enum AS ENUM ('Laki-laki', 'Perempuan');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE SEQUENCE IF NOT EXISTS pelamar_keluarga_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS pelamar_keluarga (
    id bigint NOT NULL DEFAULT nextval('pelamar_keluarga_id_seq'::regclass),
    id_keluarga VARCHAR NOT NULL,
    id_pelamar VARCHAR NOT NULL,
    nama_istri_suami VARCHAR(100),
    jenis_kelamin jenis_kelamin_enum NOT NULL,
    tempat_lahir VARCHAR(20),
    tgl_lahir DATE,
    pendidikan_terakhir VARCHAR(50),
    pekerjaan_skrg VARCHAR(50),
    alamat_rumah VARCHAR(50),
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT pelamar_keluarga_pk PRIMARY KEY (id_keluarga),
    CONSTRAINT fk_pelamar_keluarga FOREIGN KEY (id_pelamar)REFERENCES data_pelamar(id_data_pelamar)
    );