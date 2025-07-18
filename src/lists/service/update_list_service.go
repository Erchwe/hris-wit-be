package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *ListService) UpdateList(
	ctx context.Context,
	request payload.UpdateListPayload,
	user sqlc.GetUserBackofficeRow,
) error {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
			}
		}
	}()

	entity := request.ToEntity(user)
	log.FromCtx(ctx).Info(fmt.Sprintf("Attempting to update list: ListID=%s, ListName=%v, ListOrder=%v, UpdatedBy=%v",
		entity.ListID, entity.ListName.String, entity.ListOrder.Int64, entity.UpdatedBy.String))

	_, err = q.UpdateList(ctx, request.ToEntity(user))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update list")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	log.FromCtx(ctx).Info(fmt.Sprintf("List %s updated successfully with order %d", entity.ListID, entity.ListOrder.Int64))
	return nil
}
