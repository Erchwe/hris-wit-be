package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *VendorService) CreateVendor(ctx context.Context, p payload.CreateVendorPayload, user sqlc.GetUserBackofficeRow) (sqlc.Vendor, error) {
	// Validasi payload
	if err := p.Validate(); err != nil {
		log.FromCtx(ctx).Error(err, "validation failed")
		return sqlc.Vendor{}, err
	}

	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin transaction")
		return sqlc.Vendor{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.FromCtx(ctx).Error(rollbackErr, "failed to rollback")
			}
		}
	}()

	// Konversi payload ke entity
	data := p.ToEntity(user)
	fmt.Println(data)

	result, err := q.CreateVendor(ctx, data)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to insert vendor")
		return sqlc.Vendor{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "failed to commit")
		return sqlc.Vendor{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}

func (s *VendorService) CreateVendorTransactional(
	ctx context.Context,
	q *sqlc.Queries,
	param sqlc.CreateVendorParams,
) (sqlc.Vendor, error) {
	return q.CreateVendor(ctx, param)
}
