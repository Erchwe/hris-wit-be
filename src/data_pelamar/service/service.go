package service

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type DataPelamarService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewDataPelamarService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *DataPelamarService {
	return &DataPelamarService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
