package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

type VendorService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewVendorService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *VendorService {
	return &VendorService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}

func (s *VendorService) GetListVendors(ctx context.Context) ([]sqlc.Vendor, error) {
	q := sqlc.New(s.mainDB)
	result, err := q.ListVendors(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get list vendor")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}
	return result, nil
}

func (s *VendorService) UpdateVendorTransactional(
	ctx context.Context,
	q *sqlc.Queries,
	param sqlc.UpdateVendorParams,
) (sqlc.Vendor, error) {
	return q.UpdateVendor(ctx, param)
}
