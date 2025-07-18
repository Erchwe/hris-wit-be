package application

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/client/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteClient(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {

	svc := service.NewClientService(s.GetDB(), cfg)

	client := e.Group("/client")

	client.GET("", getListClient(svc))
	client.GET("/:id", getClientByID(svc))
	client.POST("", insertClient(svc, cfg))
	client.DELETE("/:id", deleteClient(svc))
	client.PUT("/:id", updateClient(svc, cfg))
}

func insertClient(svc *service.ClientService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.InsertClientPayload

		// Parse the request body
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		// Validate the request
		if err := request.Validate(); err != nil {
			return err
		}

		// Prepare user info (this could be coming from authentication, for example)
		user := sqlc.GetUserBackofficeRow{}

		// Insert the client into the database and get the full client data
		insertedClient, err := svc.InsertClient(ctx.Request().Context(), request, user, cfg)
		if err != nil {
			return err
		}

		// Return the full client data, including the client_id, to the frontend
		return httpservice.ResponseData(ctx, insertedClient, nil)
	}
}

func getListClient(svc *service.ClientService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListClient(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getClientByID(svc *service.ClientService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetClientByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteClient(svc *service.ClientService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		err := svc.DeleteClient(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "client berhasil dihapus"}, nil)
	}
}

func updateClient(svc *service.ClientService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		var request payload.UpdateClientPayload

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

		err := svc.UpdateClient(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "client berhasil diupdate"}, nil)
	}
}
