package service

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type DataPegawaiService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewDataPegawaiService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *DataPegawaiService {
	return &DataPegawaiService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
