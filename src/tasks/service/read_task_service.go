package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *TaskService) GetTaskByID(ctx context.Context, taskID string) (sqlc.Task, error) {
	q := sqlc.New(s.mainDB)

	task, err := q.GetTaskByID(ctx, taskID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get task by ID")
		return sqlc.Task{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return task, nil
}

func (s *TaskService) GetListTasks(ctx context.Context) ([]sqlc.Task, error) {
	q := sqlc.New(s.mainDB)

	tasks, err := q.ListTasks(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list of tasks")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return tasks, nil
}
