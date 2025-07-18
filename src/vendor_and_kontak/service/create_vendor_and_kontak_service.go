package service

import (
	"context"
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *VendorAndKontakService) CreateVendorWithKontak(
	ctx context.Context,
	request payload.CreateVendorWithKontakPayload,
	user sqlc.GetUserBackofficeRow,
	cfg config.KVStore,
) (sqlc.Vendor, []sqlc.KontakVendor, error) {
	if err := request.Validate(); err != nil {
		return sqlc.Vendor{}, nil, err
	}

	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return sqlc.Vendor{}, nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	vendorParam, kontakParams := request.ToEntities(cfg, user)

	vendor, err := s.vendorSvc.CreateVendorTransactional(ctx, q, vendorParam)
	if err != nil {
		_ = tx.Rollback()
		log.FromCtx(ctx).Error(err, "failed to create vendor")
		return sqlc.Vendor{}, nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	var kontakResult []sqlc.KontakVendor
	for _, kp := range kontakParams {
		res, err := q.CreateKontakVendor(ctx, kp)
		if err != nil {
			_ = tx.Rollback()
			log.FromCtx(ctx).Error(err, "failed to insert kontak vendor")
			return sqlc.Vendor{}, nil, errors.WithStack(httpservice.ErrUnknownSource)
		}
		kontakResult = append(kontakResult, res)
	}

	if err := tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "commit failed")
		return sqlc.Vendor{}, nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return vendor, kontakResult, nil
}
