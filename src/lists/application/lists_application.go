package application

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/lists/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteList(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewListService(s.GetDB(), cfg)

	project := e.Group("/project/:projectId")

	// List routes
	project.GET("/lists", getList(svc))
	project.GET("/lists/:id", getListByID(svc))
	project.POST("/lists", insertList(svc, cfg))
	project.PUT("/lists/:id", updateList(svc, cfg))
	project.DELETE("/lists/:id", deleteList(svc))
}

func insertList(svc *service.ListService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		projectId := ctx.Param("projectId")
		var request payload.InsertListPayload
		if err := ctx.Bind(&request); err != nil {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		request.ProjectID = projectId

		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		err := svc.InsertList(ctx.Request().Context(), request, user, cfg)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "list berhasil ditambahkan"}, nil)
	}
}

func getList(svc *service.ListService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		projectId := ctx.Param("projectId")
		result, err := svc.GetListByProjectID(ctx.Request().Context(), projectId)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListByID(svc *service.ListService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		projectId := ctx.Param("projectId")
		listId := ctx.Param("id")

		if projectId == "" || listId == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		requestContext := ctx.Request().Context()

		result, err := svc.GetListByProjectIDAndListID(requestContext, projectId, listId)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateList(svc *service.ListService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		projectId := ctx.Param("projectId")
		listId := ctx.Param("id")
		if listId == "" || projectId == "" {
			log.FromCtx(ctx.Request().Context()).Error(nil, fmt.Sprintf("Missing listId or projectId: listID=%s, projectID=%s", listId, projectId))
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		var request payload.UpdateListPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "Failed to bind UpdateListPayload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		request.ListID = listId

		log.FromCtx(ctx.Request().Context()).Info(fmt.Sprintf("Received UpdateListPayload in handler: ListID=%s, ListName=%s, ListOrder=%d, UpdatedBy=%s",
			request.ListID, request.ListName, request.ListOrder, request.UpdatedBy))

		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		err := svc.UpdateList(ctx.Request().Context(), request, user)
		if err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "Failed to update list in service")
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "list berhasil diupdate"}, nil)
	}
}

func deleteList(svc *service.ListService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		listId := ctx.Param("id")
		deletedBy := "admin"

		err := svc.SoftDeleteListByListID(ctx.Request().Context(), listId, deletedBy)
		if err != nil {
			return err
		}

		return ctx.JSON(200, map[string]string{
			"message": "List deleted successfully",
		})
	}
}

func getListsByProjectID(svc *service.ListService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		projectId := ctx.Param("projectId")
		if projectId == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		lists, err := svc.GetListByProjectID(ctx.Request().Context(), projectId)
		if err != nil {
			return err
		}

		if len(lists) == 0 {
			return httpservice.ResponseData(ctx, []interface{}{}, nil)
		}

		return httpservice.ResponseData(ctx, lists, nil)
	}
}
