ALTER TABLE user_backoffice
    ADD COLUMN id_pegawai varchar;

ALTER TABLE user_backoffice
    ADD CONSTRAINT fk_data_pegawai
    FOREIGN KEY (id_pegawai)
    REFERENCES data_pegawai(id_data_pegawai)
    ON DELETE SET NULL;
