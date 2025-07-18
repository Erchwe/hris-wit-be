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

func (s *LowonganPekerjaanService) InsertLowonganPekerjaan(ctx context.Context, request payload.InsertLowonganPekerjaanPayload) (lowonganPekerjaan sqlc.LowonganPekerjaan, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
				err = errors.Wrap(httpservice.ErrUnknownSource, "rollback failed")
			}
		}
	}()

	entity, err := request.ToEntity(s.cfg)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to convert payload to entity")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}
	lowonganPekerjaan, err = q.CreateLowonganPekerjaan(ctx, entity)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to insert lowongan pekerjaan")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	// Commit transaksi
	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(httpservice.ErrUnknownSource)

		return
	}

	return
}
