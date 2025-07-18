package service

import (
	"database/sql"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type BrandService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewBrandService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *BrandService {
	return &BrandService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
