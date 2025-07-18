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

// UpdateTask updates an existing task in the database.
func (s *TaskService) UpdateTask(
	ctx context.Context,
	request payload.UpdateTaskPayload,
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

	_, err = q.UpdateTask(ctx, request.ToEntity(user))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to update task")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	return nil
}
