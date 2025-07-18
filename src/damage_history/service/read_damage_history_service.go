package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *DamageHistoryService) GetDamageHistoryByID(ctx context.Context, damageHistoryID string) (sqlc.DamageHistory, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.GetDamageHistoryByID(ctx, damageHistoryID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get damage history by ID")
		return sqlc.DamageHistory{}, errors.WithStack(httpservice.ErrNoResultData)
	}

	return result, nil
}

func (s *DamageHistoryService) GetListDamageHistory(ctx context.Context) ([]sqlc.DamageHistory, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.ListDamageHistory(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to list damage history")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
