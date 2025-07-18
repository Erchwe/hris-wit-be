CREATE SEQUENCE IF NOT EXISTS attachment_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS attachments (
    id BIGINT NOT NULL DEFAULT nextval('attachment_id_seq'::regclass),
    attach_id VARCHAR NOT NULL,
    task_id VARCHAR NOT NULL,
    attach_name VARCHAR(255),
    attach_url TEXT,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(20) NOT NULL,
    updated_at TIMESTAMP,
    updated_by VARCHAR(20),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(20),
    CONSTRAINT attachments_pk PRIMARY KEY (attach_id),
    CONSTRAINT attachments_task_fk FOREIGN KEY (task_id) REFERENCES tasks(task_id)
);
