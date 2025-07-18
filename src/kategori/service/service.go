package service

import (
	"database/sql"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type KategoriService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewKategoriService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *KategoriService {
	return &KategoriService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
