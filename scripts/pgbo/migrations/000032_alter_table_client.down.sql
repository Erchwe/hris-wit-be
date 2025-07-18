DROP SEQUENCE IF EXISTS client_id_seq;

ALTER TABLE client
    ALTER COLUMN client_id SET DEFAULT nextval('client_id_seq'::regclass);
