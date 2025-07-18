package httpservice

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type Service struct {
	mainDB *sql.DB
	Cfg    config.KVStore
}

func NewService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *Service {
	return &Service{
		mainDB: mainDB,
		Cfg:    cfg,
	}
}

func (s *Service) GetDB() *sql.DB {
	return s.mainDB
}

func (s *Service) GetServiceHealth(_ context.Context) error {
	// do health check logic here
	return nil
}

// In your httpservice package, define a custom error for conflict
var ErrConflict = errors.New("conflict: the resource already exists")
