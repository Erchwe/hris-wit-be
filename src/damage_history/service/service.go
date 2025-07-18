package service

import (
	"database/sql"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type DamageHistoryService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewDamageHistoryService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *DamageHistoryService {
	return &DamageHistoryService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
