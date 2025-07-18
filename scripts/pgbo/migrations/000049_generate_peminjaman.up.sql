CREATE TYPE status_peminjaman_enum AS ENUM (
  'Menunggu Persetujuan',
  'Sedang Dipinjam',
  'Tidak Dipinjam'
);

CREATE SEQUENCE IF NOT EXISTS peminjaman_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS peminjaman (
    id BIGINT NOT NULL DEFAULT nextval('peminjaman_id_seq'::regclass),
    user_id VARCHAR NOT NULL,
    peminjaman_id VARCHAR NOT NULL,
    inventaris_id VARCHAR NOT NULL,
    tgl_pinjam DATE NOT NULL,
    tgl_kembali DATE NOT NULL,
    status_peminjaman status_peminjaman_enum NOT NULL DEFAULT 'Tidak Dipinjam',
    notes TEXT,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT peminjaman_pk PRIMARY KEY (peminjaman_id),
    CONSTRAINT inventaris_peminjaman_fk FOREIGN KEY (inventaris_id) REFERENCES inventaris(inventaris_id)
)
