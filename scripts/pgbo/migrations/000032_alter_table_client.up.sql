CREATE SEQUENCE IF NOT EXISTS client_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER TABLE client
    ALTER COLUMN client_id SET DEFAULT 'CL-' || LPAD(nextval('client_id_seq')::TEXT, 3, '0');
