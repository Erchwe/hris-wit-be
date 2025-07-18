package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *KategoriService) DeleteKategori(
	ctx context.Context,
	kategoriID string,
	deletedBy string,
) error {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.FromCtx(ctx).Error(rollbackErr, "rollback failed", err)
			}
		}
	}()

	err = q.SoftDeleteKategori(ctx, sqlc.SoftDeleteKategoriParams{
		KategoriID: kategoriID,
		DeletedBy:  sql.NullString{String: deletedBy, Valid: true},
	})

	if err != nil {
		log.FromCtx(ctx).Error(err, "soft delete kategori failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "commit failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}
