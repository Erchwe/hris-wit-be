DROP SEQUENCE IF EXISTS project_id_seq;

DROP SEQUENCE IF EXISTS id_seq;

ALTER TABLE projects
    ALTER COLUMN project_id SET DEFAULT nextval('project_id_seq'::regclass);

ALTER TABLE projects
    ALTER COLUMN id DROP DEFAULT;
