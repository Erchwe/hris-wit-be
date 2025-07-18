CREATE SEQUENCE IF NOT EXISTS pelamar_pendidikan_non_formal_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS pelamar_pendidikan_non_formal(
                                                            id bigint NOT NULL DEFAULT nextval('pelamar_pendidikan_non_formal_id_seq'::regclass),
    id_pddk_non_formal VARCHAR NOT NULL,
    id_pelamar VARCHAR NOT NULL,
    institusi VARCHAR(100) NOT NULL,
    jenis_pendidikan VARCHAR(20) NOT NULL,
    kota VARCHAR(20) NOT NULL,
    tgl_lulus DATE,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT pelamar_pendidikan_non_formal_pk PRIMARY KEY (id_pddk_non_formal),
    CONSTRAINT fk_pelamar_pendidikan_non_formal FOREIGN KEY (id_pelamar)REFERENCES data_pelamar(id_data_pelamar)
    );