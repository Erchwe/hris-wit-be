DO $$ BEGIN
    CREATE TYPE project_role_enum AS ENUM ('assignee', 'supervisor');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE SEQUENCE IF NOT EXISTS project_members_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS project_members (
    id BIGINT NOT NULL DEFAULT nextval('project_members_id_seq'::regclass),
    project_memb_id VARCHAR NOT NULL,
    id_data_pegawai VARCHAR NOT NULL,
    project_id VARCHAR NOT NULL,
    project_role project_role_enum NOT NULL,
    CONSTRAINT project_members_pk PRIMARY KEY (project_memb_id),
    CONSTRAINT project_members_project_fk FOREIGN KEY (project_id) REFERENCES projects(project_id),
    CONSTRAINT project_members_pegawai_fk FOREIGN KEY (id_data_pegawai) REFERENCES data_pegawai(id_data_pegawai)
);
