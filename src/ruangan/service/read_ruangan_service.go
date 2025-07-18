package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *RuanganService) GetRuanganByID(ctx context.Context, ruanganID string) (sqlc.Ruangan, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.GetRuanganByID(ctx, ruanganID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get ruangan by ID")
		return sqlc.Ruangan{}, errors.WithStack(httpservice.ErrNoResultData)
	}

	return result, nil
}

func (s *RuanganService) GetListRuangan(ctx context.Context) ([]sqlc.Ruangan, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.ListRuangan(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list ruangan")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
