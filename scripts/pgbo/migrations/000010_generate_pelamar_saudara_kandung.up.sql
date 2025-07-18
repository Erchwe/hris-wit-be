DO $$ BEGIN
CREATE TYPE jenis_kelamin_enum AS ENUM ('Laki-laki', 'Perempuan');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE SEQUENCE IF NOT EXISTS pelamar_saudara_kandung_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS pelamar_saudara_kandung (
                                                       id bigint NOT NULL DEFAULT nextval('pelamar_saudara_kandung_id_seq'::regclass),
    id_saudara VARCHAR NOT NULL,
    id_pelamar VARCHAR NOT NULL,
    nama VARCHAR(256) NOT NULL,
    jenis_kelamin jenis_kelamin_enum NOT NULL,
    tempat_lahir VARCHAR(20),
    pendidikan_pekerjaan VARCHAR(20),
    tgl_lahir DATE,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT pelamar_saudara_kandung_pk PRIMARY KEY (id_saudara),
    CONSTRAINT fk_pelamar_saudara_kandung FOREIGN KEY (id_pelamar)REFERENCES data_pelamar(id_data_pelamar)
    );