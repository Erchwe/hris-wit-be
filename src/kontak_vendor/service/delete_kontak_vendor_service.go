package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *KontakVendorService) DeleteKontakVendor(
	ctx context.Context,
	kontakVendorID string,
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
			if rerr := tx.Rollback(); rerr != nil {
				log.FromCtx(ctx).Error(rerr, "rollback failed", err)
			}
		}
	}()

	err = q.SoftDeleteKontakVendor(ctx, sqlc.SoftDeleteKontakVendorParams{
		KontakVendorID: kontakVendorID,
		DeletedBy:      sql.NullString{String: deletedBy, Valid: true},
	})

	if err != nil {
		log.FromCtx(ctx).Error(err, "soft delete kontak vendor failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "commit failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}
