package service

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *LowonganPekerjaanService) UpdateLowonganPekerjaan(ctx context.Context, request sqlc.UpdateLowonganPekerjaanParams) (lowongan sqlc.LowonganPekerjaan, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return lowongan, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)
	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
			}
		}
	}()

	lowongan, err = q.UpdateLowonganPekerjaan(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update lowongan pekerjaan")
		return lowongan, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return lowongan, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return lowongan, nil
}
