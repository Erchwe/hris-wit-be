package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *ListService) GetListByProjectID(ctx context.Context, projectID string) ([]sqlc.List, error) {
	q := sqlc.New(s.mainDB)

	lists, err := q.ListListsByProjectID(ctx, projectID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get lists by project ID")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return lists, nil
}

func (s *ListService) GetListByProjectIDAndListID(ctx context.Context, projectID string, listID string) (sqlc.List, error) {
	q := sqlc.New(s.mainDB)

	params := sqlc.GetListByProjectIDAndListIDParams{
		ProjectID: projectID,
		ListID:    listID,
	}

	list, err := q.GetListByProjectIDAndListID(ctx, params)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list by project ID and list ID")
		return sqlc.List{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return list, nil
}

func (s *ListService) GetListByID(ctx context.Context, listID string) (sqlc.List, error) {
	q := sqlc.New(s.mainDB)

	list, err := q.GetListByID(ctx, listID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list by ID")
		return sqlc.List{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return list, nil
}

func (s *ListService) GetList(ctx context.Context) ([]sqlc.List, error) {
	q := sqlc.New(s.mainDB)

	lists, err := q.ListLists(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list of lists")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return lists, nil
}
