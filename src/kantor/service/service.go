package service

import (
	"database/sql"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type KantorService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewKantorService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *KantorService {
	return &KantorService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
