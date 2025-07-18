package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *KategoriService) GetKategoriByID(ctx context.Context, kategoriID string) (sqlc.Kategori, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.GetKategoriByID(ctx, kategoriID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get kategori by ID")
		return sqlc.Kategori{}, errors.WithStack(httpservice.ErrNoResultData)
	}

	return result, nil
}

func (s *KategoriService) GetListKategori(ctx context.Context) ([]sqlc.Kategori, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.ListKategori(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list kategori")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
