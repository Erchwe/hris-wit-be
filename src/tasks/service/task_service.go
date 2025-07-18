package service

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type TaskService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewTaskService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *TaskService {
	return &TaskService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
