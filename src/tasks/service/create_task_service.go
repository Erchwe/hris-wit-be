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
	"time"
)

func (s *TaskService) InsertTask(
	ctx context.Context,
	request payload.InsertTaskPayload,
	user sqlc.GetUserBackofficeRow,
	cfg config.KVStore,
) (sqlc.Task, error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to begin tx")
		return sqlc.Task{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
			}
		}
	}()

	params := request.ToEntity(cfg, user)
	params.TaskOrder = sql.NullFloat64{
		Float64: float64(time.Now().UnixMilli()),
		Valid:   true,
	}

	newTask, err := q.CreateTask(ctx, params)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to insert task")
		return sqlc.Task{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return sqlc.Task{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return newTask, nil
}
