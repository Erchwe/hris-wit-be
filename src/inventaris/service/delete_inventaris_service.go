package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *InventarisService) DeleteInventaris(
	ctx context.Context,
	inventarisID string,
	user sqlc.GetUserBackofficeRow,
) error {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollbackErr)
			}
		}
	}()

	err = q.SoftDeleteInventaris(ctx, sqlc.SoftDeleteInventarisParams{
		InventarisID: inventarisID,
		DeletedBy:    sql.NullString{String: user.CreatedBy, Valid: true},
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to delete inventaris")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}
