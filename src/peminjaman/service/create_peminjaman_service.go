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

func (s *PeminjamanService) CreatePeminjaman(
	ctx context.Context,
	request payload.CreatePeminjamanPayload,
	user sqlc.GetUserBackofficeRow,
) (result payload.PeminjamanResponse, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx for create peminjaman")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback create peminjaman", rollBackErr)
			}
		}
	}()

	entity, err := request.ToEntity(user)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to convert payload to peminjaman entity")
		return result, errors.WithStack(httpservice.ErrBadRequest)
	}

	createdPeminjaman, err := q.CreatePeminjaman(ctx, entity)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed create peminjaman: %v", err)
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	result = payload.PeminjamanResponse{
		PeminjamanID:     createdPeminjaman.PeminjamanID,
		TanggalPinjam:    createdPeminjaman.TglPinjam,
		TanggalKembali:   createdPeminjaman.TglKembali,
		StatusPeminjaman: string(createdPeminjaman.StatusPeminjaman),
		Notes:            createdPeminjaman.Notes.String,
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit create peminjaman")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
