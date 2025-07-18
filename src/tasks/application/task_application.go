package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/src/tasks/service"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

func AddRouteTask(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewTaskService(s.GetDB(), cfg)

	task := e.Group("/task")

	task.GET("", getListTask(svc))
	task.GET("/:id", getTaskByID(svc))
	task.POST("", insertTask(svc, cfg))
	task.PUT("/:id", updateTask(svc, cfg))
	task.DELETE("/:id", deleteTask(svc))
}

func insertTask(svc *service.TaskService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.InsertTaskPayload
		if err := ctx.Bind(&request); err != nil {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		newTask, err := svc.InsertTask(ctx.Request().Context(), request, user, cfg)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, newTask, nil)
	}
}

func getListTask(svc *service.TaskService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListTasks(ctx.Request().Context())
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getTaskByID(svc *service.TaskService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetTaskByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateTask(svc *service.TaskService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		var request payload.UpdateTaskPayload
		if err := ctx.Bind(&request); err != nil {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		err := svc.UpdateTask(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "task berhasil diupdate"}, nil)
	}
}

func deleteTask(svc *service.TaskService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		err := svc.DeleteTask(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "task berhasil dihapus"}, nil)
	}
}
