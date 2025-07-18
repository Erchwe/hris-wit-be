package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *VendorAndKontakService) DeleteVendorWithKontak(
	ctx context.Context,
	vendorID string,
	user sqlc.GetUserBackofficeRow,
) error {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	// Soft delete kontak vendor by vendor ID
	err = q.SoftDeleteKontakVendorByVendorID(ctx, sqlc.SoftDeleteKontakVendorByVendorIDParams{
		VendorID:  vendorID,
		DeletedBy: sql.NullString{String: user.DeletedBy.String, Valid: user.DeletedBy.Valid},
	})
	if err != nil {
		_ = tx.Rollback()
		log.FromCtx(ctx).Error(err, "failed to soft delete kontak vendor")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	// Soft delete vendor
	err = q.SoftDeleteVendor(ctx, sqlc.SoftDeleteVendorParams{
		VendorID:  vendorID,
		DeletedBy: sql.NullString{String: user.DeletedBy.String, Valid: user.DeletedBy.Valid},
	})
	if err != nil {
		_ = tx.Rollback()
		log.FromCtx(ctx).Error(err, "failed to soft delete vendor")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err := tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "commit failed")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}
