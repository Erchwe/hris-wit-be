package service

import (
	"database/sql"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type KontakVendorService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewKontakVendorService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *KontakVendorService {
	return &KontakVendorService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
