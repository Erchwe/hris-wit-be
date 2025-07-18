package service

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

// deleteLowonganPekerjaan
func (s *LowonganPekerjaanService) DeleteLowonganPekerjaan(
	ctx context.Context,
	id string,
	// userData sqlc.GetUserBackofficeByEmailRow,
) (err error) {
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
				err = errors.WithStack(httpservice.ErrUnknownSource)

				return
			}
		}
	}()

	err = q.SoftDeleteLowonganPekerjaan(ctx, sqlc.SoftDeleteLowonganPekerjaanParams{
		IDLowonganPekerjaan: id,
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to delete lowongan pekerjaan")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}
