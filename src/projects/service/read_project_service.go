package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *ProjectService) GetProjectByID(ctx context.Context, id string) (sqlc.Project, error) {
	q := sqlc.New(s.mainDB)

	project, err := q.GetProjectByID(ctx, id)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get project by ID")
		return sqlc.Project{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return project, nil
}

func (s *ProjectService) GetListProjects(ctx context.Context) ([]sqlc.Project, error) {
	q := sqlc.New(s.mainDB)

	projects, err := q.ListProjects(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list of projects")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return projects, nil
}
