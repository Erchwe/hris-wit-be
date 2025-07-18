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

func (s *DataPegawaiService) InsertDataPegawai(ctx context.Context, request payload.InsertDataPegawaiPayload) (pegawai sqlc.DataPegawai, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
			}
		}
	}()

	params := request.ToEntity(s.cfg)

	pegawai, err = q.CreateDataPegawai(ctx, params)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to insert data pegawai")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		err = errors.WithStack(httpservice.ErrUnknownSource)
		return
	}

	return
}
