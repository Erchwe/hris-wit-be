package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/src/usage_history/service"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteUsageHistory(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewUsageHistoryService(s.GetDB(), cfg)

	route := e.Group("/usage-history")
	route.POST("", insertUsageHistory(svc))
	route.GET("", getListUsageHistory(svc))
	route.GET("/:id", getUsageHistoryByID(svc))
	route.PUT("/:id", updateUsageHistory(svc))
	route.DELETE("/:id", deleteUsageHistory(svc))
}

func insertUsageHistory(svc *service.UsageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreateUsageHistoryPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.CreateUsageHistory(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListUsageHistory(svc *service.UsageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListUsageHistory(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getUsageHistoryByID(svc *service.UsageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetUsageHistoryByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateUsageHistory(svc *service.UsageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.UpdateUsageHistoryPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}

		result, err := svc.UpdateUsageHistory(ctx.Request().Context(), request)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteUsageHistory(svc *service.UsageHistoryService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		user := sqlc.GetUserBackofficeRow{}
		err := svc.DeleteUsageHistory(ctx.Request().Context(), id, user.CreatedBy)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "usage history deleted"}, nil)
	}
}
