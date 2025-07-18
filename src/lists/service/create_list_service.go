package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
	"time"
)

func (s *ListService) InsertList(
	ctx context.Context,
	request payload.InsertListPayload,
	user sqlc.GetUserBackofficeRow,
	cfg config.KVStore,
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

	if request.ListOrder == 0 {
		request.ListOrder = time.Now().UnixMilli()
	}

	entity := request.ToEntity(cfg, user)
	log.FromCtx(ctx).Info(fmt.Sprintf("Attempting to insert list: ProjectID=%s, ListName=%v, ListOrder=%v, CreatedBy=%v",
		entity.ProjectID, entity.ListName.String, entity.ListOrder.Int64, entity.CreatedBy))

	_, err = q.CreateList(ctx, request.ToEntity(cfg, user))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to insert list")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}
	return nil
}
