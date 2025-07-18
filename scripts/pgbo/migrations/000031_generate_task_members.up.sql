CREATE SEQUENCE IF NOT EXISTS task_members_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS task_members (
    id BIGINT NOT NULL DEFAULT nextval('task_members_id_seq'::regclass),
    task_memb_id VARCHAR NOT NULL,
    project_memb_id VARCHAR NOT NULL,
    task_id VARCHAR NOT NULL,
    CONSTRAINT task_members_pk PRIMARY KEY (task_memb_id),
    CONSTRAINT task_members_project_members_fk FOREIGN KEY (project_memb_id) REFERENCES project_members(project_memb_id),
    CONSTRAINT task_members_task_fk FOREIGN KEY (task_id) REFERENCES tasks(task_id)
);
