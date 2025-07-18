package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/kantor/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteKantor(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewKantorService(s.GetDB(), cfg)

	r := e.Group("/kantor")
	r.POST("/insert", insertKantor(svc))
	r.GET("", getListKantor(svc))
	r.GET("/:id", getKantorByID(svc))
	r.PUT("/update:id", updateKantor(svc))
	r.DELETE("/delete:id", deleteKantor(svc))
}

func insertKantor(svc *service.KantorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreateKantorPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.CreateKantor(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListKantor(svc *service.KantorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListKantor(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getKantorByID(svc *service.KantorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		result, err := svc.GetKantorByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateKantor(svc *service.KantorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.UpdateKantorPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}

		result, err := svc.UpdateKantor(ctx.Request().Context(), request)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteKantor(svc *service.KantorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		user := sqlc.GetUserBackofficeRow{}
		err := svc.DeleteKantor(ctx.Request().Context(), id, user.CreatedBy)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, map[string]string{"message": "kantor deleted"}, nil)
	}
}
