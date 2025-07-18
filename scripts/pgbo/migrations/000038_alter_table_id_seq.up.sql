ALTER TABLE tasks
    ALTER COLUMN task_id SET DEFAULT 'TA-' || LPAD(nextval('tasks_id_seq')::TEXT, 3, '0');