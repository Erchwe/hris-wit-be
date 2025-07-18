package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *InventarisService) UpdateInventaris(
	ctx context.Context,
	request payload.UpdateInventarisPayload,
	user sqlc.GetUserBackofficeRow,
	cfg config.KVStore,
) (result sqlc.Inventari, err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	entity := request.ToEntity()

	result, err = q.UpdateInventaris(ctx, entity)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update inventaris")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
