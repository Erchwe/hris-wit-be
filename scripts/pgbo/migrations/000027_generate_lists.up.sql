CREATE SEQUENCE IF NOT EXISTS list_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS lists (
    id BIGINT NOT NULL DEFAULT nextval('list_id_seq'::regclass),
    list_id VARCHAR NOT NULL,
    project_id VARCHAR NOT NULL,
    list_name VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(20) NOT NULL,
    updated_at TIMESTAMP,
    updated_by VARCHAR(20),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(20),
    CONSTRAINT lists_pk PRIMARY KEY (list_id),
    CONSTRAINT lists_project_fk FOREIGN KEY (project_id) REFERENCES projects(project_id)
);
