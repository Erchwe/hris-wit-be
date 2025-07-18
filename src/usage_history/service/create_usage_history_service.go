package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *UsageHistoryService) CreateUsageHistory(
	ctx context.Context,
	p payload.CreateUsageHistoryPayload,
	user sqlc.GetUserBackofficeRow,
) (sqlc.UsageHistory, error) {
	if err := p.Validate(); err != nil {
		log.FromCtx(ctx).Error(err, "validation error")
		return sqlc.UsageHistory{}, err
	}

	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return sqlc.UsageHistory{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.FromCtx(ctx).Error(rerr, "rollback failed", err)
			}
		}
	}()

	arg, err := p.ToEntity(s.cfg, user)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to convert payload to entity")
		return sqlc.UsageHistory{}, errors.WithStack(httpservice.ErrBadRequest)
	}

	result, err := q.CreateUsageHistory(ctx, arg)
	if err != nil {
		log.FromCtx(ctx).Error(err, "create usage history failed")
		return sqlc.UsageHistory{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "commit failed")
		return sqlc.UsageHistory{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
