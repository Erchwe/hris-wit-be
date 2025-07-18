package service

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *DataPelamarService) UpdateStatusPelamar(ctx context.Context, request sqlc.UpdatePelamarParams) (pelamar sqlc.DataPelamar, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return pelamar, errors.WithStack(httpservice.ErrUnknownSource)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(rollBackErr, "error on rollback after an error", "original_error", err)
			}
		} else {
			err = tx.Commit()
			if err != nil {
				log.FromCtx(ctx).Error(err, "error on commit")
				err = errors.WithStack(httpservice.ErrUnknownSource)
			}
		}
	}()

	q := sqlc.New(s.mainDB).WithTx(tx)

	pelamar, err = q.UpdatePelamar(ctx, request)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update pelamar status")
		if errors.Is(err, sql.ErrNoRows) {
			return pelamar, errors.WithStack(httpservice.ErrUnknownSource)
		}
		return pelamar, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return pelamar, nil
}
