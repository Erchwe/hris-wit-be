DO $$ BEGIN
CREATE TYPE jenis_kelamin_enum AS ENUM ('Laki-laki', 'Perempuan');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
CREATE TYPE status_enum AS ENUM (
    'New',
    'Short list',
    'HR Interview',
    'User Interview',
    'Refference Checking',
    'Offering',
    'Psikotest',
    'Hired',
    'Rejected'
);
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;


CREATE SEQUENCE IF NOT EXISTS data_pegawai_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS data_pegawai (
    id bigint NOT NULL DEFAULT nextval('data_pegawai_id_seq'::regclass),
    id_data_pegawai VARCHAR NOT NULL,
    employee_number VARCHAR UNIQUE NOT NULL,
    divisi VARCHAR(20),
    nama_lengkap VARCHAR(256) NOT NULL,
    tempat_lahir VARCHAR(20),
    tgl_lahir DATE,
    jenis_kelamin jenis_kelamin_enum NOT NULL,
    kewarganegaraan VARCHAR(50),
    phone VARCHAR(15),
    mobile VARCHAR(15),
    agama VARCHAR(10),
    gol_darah VARCHAR(2),
    gaji MONEY,
    status_menikah BOOLEAN,
    no_ktp VARCHAR(16) UNIQUE,
    no_npwp VARCHAR(16) UNIQUE,
    status status_enum,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT data_pegawai_pk PRIMARY KEY (id_data_pegawai)
    );

INSERT INTO data_pegawai (
    id_data_pegawai,
    employee_number,
    nama_lengkap,
    jenis_kelamin,
    created_at,
    created_by
) VALUES (
'pegawai-001',
'EMP001',
'Admin',
'Laki-laki',
(now() at time zone 'UTC')::TIMESTAMP,
'system'
);
