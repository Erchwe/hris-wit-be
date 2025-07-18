DO $$ BEGIN
    CREATE TYPE project_status_enum AS ENUM ('open', 'on progress', 'done');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE SEQUENCE IF NOT EXISTS project_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS projects (
    id BIGINT NOT NULL DEFAULT nextval('project_id_seq'::regclass),
    project_id VARCHAR NOT NULL,
    client_id VARCHAR NOT NULL,
    project_name VARCHAR(255),
    project_desc TEXT,
    project_status project_status_enum,
    project_priority VARCHAR(50),
    project_color CHAR(7),
    start_date TIMESTAMP,
    due_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(20) NOT NULL,
    updated_at TIMESTAMP,
    updated_by VARCHAR(20),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(20),
    CONSTRAINT projects_pk PRIMARY KEY (project_id),
    CONSTRAINT projects_client_fk FOREIGN KEY (client_id) REFERENCES client(client_id)
);
