package application

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/projects/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteProject(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewProjectService(s.GetDB(), cfg)

	project := e.Group("/project")

	project.GET("", getListProject(svc))
	project.GET("/:id", getProjectByID(svc))
	project.POST("", insertProject(svc, cfg))
	project.DELETE("/:id", deleteProject(svc))
	project.PUT("/:id", updateProject(svc, cfg))
}

func insertProject(svc *service.ProjectService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.InsertProjectPayload

		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		err := svc.InsertProject(ctx.Request().Context(), request, user, cfg)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "project berhasil ditambahkan"}, nil)
	}
}

func getListProject(svc *service.ProjectService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.FromCtx(ctx.Request().Context()).Info("GET /project route triggered")

		result, err := svc.GetListProjects(ctx.Request().Context())
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getProjectByID(svc *service.ProjectService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetProjectByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteProject(svc *service.ProjectService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		err := svc.DeleteProject(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "project berhasil dihapus"}, nil)
	}
}

func updateProject(svc *service.ProjectService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		var request payload.UpdateProjectPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{
			ID:        10001,
			CreatedBy: "John-Doe",
			UpdatedBy: sql.NullString{String: "dummy-user", Valid: true},
		}

		err := svc.UpdateProject(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "project berhasil diupdate"}, nil)
	}
}
