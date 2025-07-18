CREATE SEQUENCE IF NOT EXISTS usage_history_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

CREATE TABLE IF NOT EXISTS usage_history (
    id BIGINT NOT NULL DEFAULT nextval('usage_history_id_seq'::regclass),
    usage_history_id VARCHAR NOT NULL,
    inventaris_id VARCHAR NOT NULL,
    old_room_id VARCHAR NOT NULL,
    new_room_id VARCHAR NOT NULL,
    old_user_id VARCHAR NOT NULL,
    new_user_id VARCHAR NOT NULL,
    moved_at TIMESTAMP,
    created_at timestamp without time zone NOT NULL,
    created_by character varying NOT NULL,
    updated_at timestamp without time zone,
    updated_by character varying,
    deleted_at timestamp without time zone,
    deleted_by character varying,
    CONSTRAINT usage_history_pk PRIMARY KEY (usage_history_id),
    CONSTRAINT inventaris_usage_history_fk FOREIGN KEY (inventaris_id) REFERENCES inventaris(inventaris_id),
    CONSTRAINT old_room_usage_history_fk FOREIGN KEY (old_room_id) REFERENCES ruangan(ruangan_id),
    CONSTRAINT new_room_usage_history_fk FOREIGN KEY (new_room_id) REFERENCES ruangan(ruangan_id),
    CONSTRAINT old_user_usage_history_fk FOREIGN KEY (old_user_id) REFERENCES data_pegawai(id_data_pegawai),
    CONSTRAINT new_user_usage_history_fk FOREIGN KEY (new_user_id) REFERENCES data_pegawai(id_data_pegawai)
);