package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *LowonganPekerjaanService) GetLowonganPekerjaan(ctx context.Context, idLowongan string) (lowongan sqlc.LowonganPekerjaan, err error) {
	q := sqlc.New(s.mainDB)

	lowongan, err = q.GetLowonganPekerjaanByID(ctx, idLowongan)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get lowongan pekerjaan")
		return lowongan, errors.Wrap(httpservice.ErrUnknownSource, "failed to get lowongan pekerjaan")
	}

	return lowongan, nil
}

func (s *LowonganPekerjaanService) ListLowonganPekerjaan(ctx context.Context) ([]sqlc.LowonganPekerjaan, error) {
	q := sqlc.New(s.mainDB)
	result, err := q.ListLowonganPekerjaan(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get list pelamar")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}
	return result, nil
}
