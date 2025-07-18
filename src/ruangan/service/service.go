package service

import (
	"database/sql"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type RuanganService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewRuanganService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *RuanganService {
	return &RuanganService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
