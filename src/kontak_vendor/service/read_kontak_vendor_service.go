package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *KontakVendorService) GetKontakVendorByID(ctx context.Context, kontakVendorID string) (sqlc.KontakVendor, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.GetKontakVendorByID(ctx, kontakVendorID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get kontak vendor by ID")
		return sqlc.KontakVendor{}, errors.WithStack(httpservice.ErrNoResultData)
	}

	return result, nil
}

func (s *KontakVendorService) GetKontakVendorListByVendorID(ctx context.Context, vendorID string) ([]sqlc.KontakVendor, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.ListKontakVendorByVendorID(ctx, vendorID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get kontak vendor list by vendor ID")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
