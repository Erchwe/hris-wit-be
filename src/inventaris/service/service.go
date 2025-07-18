package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

type InventarisService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewInventarisService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *InventarisService {
	return &InventarisService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}

func (s *InventarisService) GetListInventaris(ctx context.Context) ([]sqlc.Inventari, error) {
	q := sqlc.New(s.mainDB)
	result, err := q.ListInventaris(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get list inventaris")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}
	return result, nil
}

func (s *InventarisService) GetListInventarisWithRelations(ctx context.Context) ([]sqlc.ListInventarisWithRelationsRow, error) {
	q := sqlc.New(s.mainDB)
	result, err := q.ListInventarisWithRelations(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get list inventaris with relation")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}
	return result, nil
}

func (s *InventarisService) GetInventarisWithRelationsByID(ctx context.Context, id string) (sqlc.GetInventarisWithRelationsByIDRow, error) {
	q := sqlc.New(s.mainDB)
	result, err := q.GetInventarisWithRelationsByID(ctx, id)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get inventaris with relations by id")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}
	return result, nil
}
