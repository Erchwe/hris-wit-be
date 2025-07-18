package service

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type AttachmentService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewAttachmentService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *AttachmentService {
	return &AttachmentService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
