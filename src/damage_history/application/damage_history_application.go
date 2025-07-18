package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/damage_history/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteDamageHistory(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewDamageHistoryService(s.GetDB(), cfg)

	d := e.Group("/damage-history")
	d.POST("", insertDamageHistory(svc))
	d.GET("", getListDamageHistory(svc))
	d.GET("/:id", getDamageHistoryByID(svc))
	d.PUT("/:id", updateDamageHistory(svc))
	d.DELETE("/:id", deleteDamageHistory(svc))
}

func insertDamageHistory(svc *service.DamageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreateDamageHistoryPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.CreateDamageHistory(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListDamageHistory(svc *service.DamageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListDamageHistory(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getDamageHistoryByID(svc *service.DamageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		result, err := svc.GetDamageHistoryByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateDamageHistory(svc *service.DamageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.UpdateDamageHistoryPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		result, err := svc.UpdateDamageHistory(ctx.Request().Context(), request)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteDamageHistory(svc *service.DamageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		user := sqlc.GetUserBackofficeRow{}
		err := svc.DeleteDamageHistory(ctx.Request().Context(), id, user.CreatedBy)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, map[string]string{"message": "damage history deleted"}, nil)
	}
}
