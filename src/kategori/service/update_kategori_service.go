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

func (s *KategoriService) UpdateKategori(ctx context.Context, p payload.UpdateKategoriPayload) (sqlc.Kategori, error) {
	if err := p.Validate(); err != nil {
		log.FromCtx(ctx).Error(err, "validation error")
		return sqlc.Kategori{}, err
	}

	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return sqlc.Kategori{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.FromCtx(ctx).Error(rollbackErr, "rollback failed", err)
			}
		}
	}()

	arg := p.ToEntity()

	result, err := q.UpdateKategori(ctx, arg)
	if err != nil {
		log.FromCtx(ctx).Error(err, "update kategori failed")
		return sqlc.Kategori{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "commit failed")
		return sqlc.Kategori{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
