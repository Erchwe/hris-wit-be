package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *ClientService) GetClientByID(ctx context.Context, id string) (sqlc.Client, error) {
	q := sqlc.New(s.mainDB)

	client, err := q.GetClientByID(ctx, id)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get client by ID")
		return sqlc.Client{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return client, nil
}

func (s *ClientService) GetListClient(ctx context.Context) ([]sqlc.Client, error) {
	q := sqlc.New(s.mainDB)
	result, err := q.ListClients(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list of clients")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}
	return result, nil
}
