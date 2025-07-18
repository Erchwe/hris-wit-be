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

func (s *InventarisService) CreateInventaris(
	ctx context.Context,
	request payload.CreateInventarisPayload,
	user sqlc.GetUserBackofficeRow,
) (result sqlc.Inventari, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
			}
		}
	}()

	entity, err := request.ToEntity(user)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed convert inventaris entity")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	result, err = q.CreateInventaris(ctx, entity)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed insert inventaris")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
