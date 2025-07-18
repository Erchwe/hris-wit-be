package service

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type ClientService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewClientService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *ClientService {
	return &ClientService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
