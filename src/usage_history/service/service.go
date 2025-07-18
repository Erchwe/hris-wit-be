package service

import (
	"database/sql"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type UsageHistoryService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewUsageHistoryService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *UsageHistoryService {
	return &UsageHistoryService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
