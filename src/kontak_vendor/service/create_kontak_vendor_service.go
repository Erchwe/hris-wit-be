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

func (s *KontakVendorService) CreateKontakVendor(
	ctx context.Context,
	p payload.CreateKontakVendorPayload,
	user sqlc.GetUserBackofficeRow,
) (sqlc.KontakVendor, error) {
	if err := p.Validate(); err != nil {
		log.FromCtx(ctx).Error(err, "validation error")
		return sqlc.KontakVendor{}, err
	}

	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return sqlc.KontakVendor{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				log.FromCtx(ctx).Error(rerr, "rollback failed", err)
			}
		}
	}()

	arg := p.ToEntity(user)

	result, err := q.CreateKontakVendor(ctx, arg)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to create kontak vendor")
		return sqlc.KontakVendor{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "commit failed")
		return sqlc.KontakVendor{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
