package service

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type ListService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewListService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *ListService {
	return &ListService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
