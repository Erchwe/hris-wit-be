package service

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type ProjectService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewProjectService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *ProjectService {
	return &ProjectService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
