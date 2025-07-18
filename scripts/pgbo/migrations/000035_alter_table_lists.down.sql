DROP SEQUENCE IF EXISTS list_id_seq;

DROP SEQUENCE IF EXISTS id_seq;

ALTER TABLE lists
    ALTER COLUMN list_id SET DEFAULT nextval('list_id_seq'::regclass);

ALTER TABLE lists
    ALTER COLUMN id DROP DEFAULT;
