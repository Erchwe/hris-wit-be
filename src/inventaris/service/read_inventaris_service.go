package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *InventarisService) GetInventarisByID(ctx context.Context, id string) (sqlc.Inventari, error) {
	q := sqlc.New(s.mainDB)

	result, err := q.GetInventarisByID(ctx, id)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get inventaris by id")
		return sqlc.Inventari{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
