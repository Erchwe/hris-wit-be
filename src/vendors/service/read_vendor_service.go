package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *VendorService) GetVendorByID(ctx context.Context, vendorID string) (sqlc.Vendor, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.GetVendorByID(ctx, vendorID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get vendor by ID")
		return sqlc.Vendor{}, errors.WithStack(httpservice.ErrNoResultData)
	}

	return result, nil
}
