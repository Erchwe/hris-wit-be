package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

// DeletePeminjaman melakukan soft delete pada peminjaman
func (s *PeminjamanService) DeletePeminjaman(ctx context.Context, peminjamanID string, user sqlc.GetUserBackofficeRow) error {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx for delete peminjaman")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback delete peminjaman", rollBackErr)
			}
		}
	}()

	err = q.SoftDeletePeminjaman(ctx, sqlc.SoftDeletePeminjamanParams{
		DeletedBy:    sql.NullString{String: user.Name.String, Valid: user.Name.Valid}, // Menggunakan user.Name
		PeminjamanID: peminjamanID,
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed soft delete peminjaman: %v", err)
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit delete peminjaman")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}

// RestorePeminjaman mengembalikan peminjaman yang sudah di-soft delete
func (s *PeminjamanService) RestorePeminjaman(ctx context.Context, peminjamanID string) error {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx for restore peminjaman")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback restore peminjaman", rollBackErr)
			}
		}
	}()

	err = q.RestorePeminjaman(ctx, peminjamanID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed restore peminjaman: %v", err)
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit restore peminjaman")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}
