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

func (s *VendorAndKontakService) UpdateVendorWithKontak(
	ctx context.Context,
	request payload.UpdateVendorWithKontakPayload,
	user sqlc.GetUserBackofficeRow,
	_ interface{}, // jika memang tidak dipakai
) (sqlc.Vendor, []sqlc.KontakVendor, error) {
	// Validasi payload
	if err := request.Validate(); err != nil {
		return sqlc.Vendor{}, nil, err
	}

	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return sqlc.Vendor{}, nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.FromCtx(ctx).Error(rollbackErr, "rollback failed", err)
			}
		}
	}()

	vendorParam, kontakParams := request.ToEntities(user)

	// Update vendor
	vendor, err := s.vendorSvc.UpdateVendorTransactional(ctx, q, vendorParam)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update vendor")
		return sqlc.Vendor{}, nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	// Soft delete kontak vendor lama
	err = q.SoftDeleteKontakVendorByVendorID(ctx, sqlc.SoftDeleteKontakVendorByVendorIDParams{
		VendorID: vendorParam.VendorID,
		DeletedBy: sql.NullString{
			String: user.UpdatedBy.String,
			Valid:  user.UpdatedBy.Valid,
		},
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to soft delete kontak vendor")
		return sqlc.Vendor{}, nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	// Insert kontak vendor baru
	var kontakResult []sqlc.KontakVendor
	for _, kp := range kontakParams {
		res, err := q.CreateKontakVendor(ctx, kp)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed to insert kontak vendor")
			return sqlc.Vendor{}, nil, errors.WithStack(httpservice.ErrUnknownSource)
		}
		kontakResult = append(kontakResult, res)
	}

	// Commit transaksi
	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "commit failed")
		return sqlc.Vendor{}, nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return vendor, kontakResult, nil
}
