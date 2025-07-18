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

func (s *DamageHistoryService) CreateDamageHistory(
	ctx context.Context,
	p payload.CreateDamageHistoryPayload,
	user sqlc.GetUserBackofficeRow,
) (sqlc.DamageHistory, error) {
	if err := p.Validate(); err != nil {
		log.FromCtx(ctx).Error(err, "validation error")
		return sqlc.DamageHistory{}, err
	}

	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return sqlc.DamageHistory{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.FromCtx(ctx).Error(rollbackErr, "rollback failed", err)
			}
		}
	}()

	arg, err := p.ToEntity(s.cfg, user)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to convert payload to entity")
		return sqlc.DamageHistory{}, errors.WithStack(httpservice.ErrBadRequest)
	}

	result, err := q.CreateDamageHistory(ctx, arg)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to create damage history")
		return sqlc.DamageHistory{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "commit failed")
		return sqlc.DamageHistory{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
