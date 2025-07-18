package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/src/ruangan/service"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteRuangan(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewRuanganService(s.GetDB(), cfg)

	r := e.Group("/ruangan")
	r.POST("/insert", insertRuangan(svc))
	r.GET("", getListRuangan(svc))
	r.GET("/:id", getRuanganByID(svc))
	r.PUT("/update:id", updateRuangan(svc))
	r.DELETE("/delete:id", deleteRuangan(svc))
}

func insertRuangan(svc *service.RuanganService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreateRuanganPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.CreateRuangan(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListRuangan(svc *service.RuanganService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListRuangan(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getRuanganByID(svc *service.RuanganService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		result, err := svc.GetRuanganByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateRuangan(svc *service.RuanganService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.UpdateRuanganPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		result, err := svc.UpdateRuangan(ctx.Request().Context(), request)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteRuangan(svc *service.RuanganService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		user := sqlc.GetUserBackofficeRow{}
		if err := svc.DeleteRuangan(ctx.Request().Context(), id, user.CreatedBy); err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, map[string]string{"message": "ruangan deleted"}, nil)
	}
}
