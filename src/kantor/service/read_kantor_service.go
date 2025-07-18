package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *KantorService) GetKantorByID(ctx context.Context, kantorID string) (sqlc.Kantor, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.GetKantorByID(ctx, kantorID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get kantor by ID")
		return sqlc.Kantor{}, errors.WithStack(httpservice.ErrNoResultData)
	}

	return result, nil
}

func (s *KantorService) GetListKantor(ctx context.Context) ([]sqlc.Kantor, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.ListKantor(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to list kantor")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
