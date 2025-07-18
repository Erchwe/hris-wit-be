DROP SEQUENCE IF EXISTS id_seq_for_tasks;

ALTER TABLE tasks
    ALTER COLUMN id DROP DEFAULT;

DROP SEQUENCE IF EXISTS tasks_id_seq;

ALTER TABLE tasks
    ALTER COLUMN project_id SET DEFAULT nextval('tasks_id_seq'::regclass);
