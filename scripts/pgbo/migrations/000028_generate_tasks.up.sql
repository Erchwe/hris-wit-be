DO $$ BEGIN
    CREATE TYPE task_type_enum AS ENUM ('task', 'bug');
    CREATE TYPE task_priority_enum AS ENUM ('low', 'medium', 'high', 'crucial');
    CREATE TYPE task_size_enum AS ENUM ('easy', 'medium', 'hard');
    CREATE TYPE task_status_enum AS ENUM ('open', 'on progress', 'done');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE SEQUENCE IF NOT EXISTS task_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS tasks (
    id BIGINT NOT NULL DEFAULT nextval('task_id_seq'::regclass),
    task_id VARCHAR NOT NULL,
    list_id VARCHAR NOT NULL,
    task_name VARCHAR(255),
    task_type task_type_enum,
    task_priority task_priority_enum,
    task_size task_size_enum,
    task_status task_status_enum,
    task_color CHAR(7),
    start_date TIMESTAMP,
    due_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(20) NOT NULL,
    updated_at TIMESTAMP,
    updated_by VARCHAR(20),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(20),
    CONSTRAINT tasks_pk PRIMARY KEY (task_id),
    CONSTRAINT tasks_list_fk FOREIGN KEY (list_id) REFERENCES lists(list_id)
);
