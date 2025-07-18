package service

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *ListService) SoftDeleteListByListID(ctx context.Context, listID, deletedBy string) error {
	q := sqlc.New(s.mainDB)

	deletedByValue := sql.NullString{
		String: deletedBy,
		Valid:  deletedBy != "",
	}

	err := q.SoftDeleteList(ctx, sqlc.SoftDeleteListParams{
		DeletedBy: deletedByValue,
		ListID:    listID,
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to soft delete list")
		return errors.Wrap(err, "failed to delete list")
	}

	return nil
}
