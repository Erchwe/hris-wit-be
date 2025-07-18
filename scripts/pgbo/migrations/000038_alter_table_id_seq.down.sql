DROP SEQUENCE IF EXISTS tasks_id_seq;

ALTER TABLE tasks
    ALTER COLUMN project_id SET DEFAULT nextval('tasks_id_seq'::regclass);