package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *DataPelamarService) GetPelamarByID(ctx context.Context, id string) (sqlc.DataPelamar, error) {
	q := sqlc.New(s.mainDB)

	pelamar, err := q.GetPelamarByID(ctx, id)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get pelamar by id")
		return sqlc.DataPelamar{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return pelamar, nil
}

// todo: bikin get by email
func (s *DataPelamarService) GetPelamarByEmail(ctx context.Context, email string) (sqlc.DataPelamar, error) {
	q := sqlc.New(s.mainDB)

	userBackoffice, err := q.GetUserBackofficeByEmail(ctx, email)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get user backoffice by email")
		return sqlc.DataPelamar{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	if userBackoffice.ID != 0 {
		return sqlc.DataPelamar{}, errors.WithStack(httpservice.ErrConflict)
	}

	pelamar, err := q.GetPelamarByID(ctx, email)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get pelamar by email")
		return sqlc.DataPelamar{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return pelamar, nil
}

func (s *DataPelamarService) ListPelamar(ctx context.Context) ([]sqlc.DataPelamar, error) {
	q := sqlc.New(s.mainDB)
	result, err := q.ListPelamar(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed get list pelamar")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}
	return result, nil
}
