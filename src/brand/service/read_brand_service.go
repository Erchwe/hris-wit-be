package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *BrandService) GetBrandByID(ctx context.Context, brandID string) (sqlc.Brand, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.GetBrandByID(ctx, brandID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get brand by ID")
		return sqlc.Brand{}, errors.WithStack(httpservice.ErrNoResultData)
	}

	return result, nil
}

func (s *BrandService) GetListBrand(ctx context.Context) ([]sqlc.Brand, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.ListBrands(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to list brand")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
