package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *DataPegawaiService) GetAllDataPegawai(ctx context.Context) ([]sqlc.DataPegawai, error) {
	q := sqlc.New(s.mainDB)
	result, err := q.ListDataPegawai(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get list pelamar")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}
	return result, nil
}
