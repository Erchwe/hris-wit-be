package service

import (
	"database/sql"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type PeminjamanService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewPeminjamanService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *PeminjamanService {
	return &PeminjamanService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
