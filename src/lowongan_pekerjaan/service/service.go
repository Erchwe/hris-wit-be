package service

import (
	"database/sql"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type LowonganPekerjaanService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewLowonganPekerjaanService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *LowonganPekerjaanService {
	return &LowonganPekerjaanService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
