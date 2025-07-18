package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/kontak_vendor/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteKontakVendor(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewKontakVendorService(s.GetDB(), cfg)

	k := e.Group("/kontak-vendor")
	k.POST("", insertKontakVendor(svc))
	k.GET("/:id", getKontakVendorByID(svc))
	k.GET("/vendor/:vendor_id", getKontakVendorListByVendorID(svc))
	k.PUT("/:id", updateKontakVendor(svc))
	k.DELETE("/:id", deleteKontakVendor(svc))
}

func insertKontakVendor(svc *service.KontakVendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreateKontakVendorPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.CreateKontakVendor(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getKontakVendorByID(svc *service.KontakVendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		result, err := svc.GetKontakVendorByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getKontakVendorListByVendorID(svc *service.KontakVendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		vendorID := ctx.Param("vendor_id")
		if vendorID == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		result, err := svc.GetKontakVendorListByVendorID(ctx.Request().Context(), vendorID)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateKontakVendor(svc *service.KontakVendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.UpdateKontakVendorPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		result, err := svc.UpdateKontakVendor(ctx.Request().Context(), request)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteKontakVendor(svc *service.KontakVendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		user := sqlc.GetUserBackofficeRow{}
		err := svc.DeleteKontakVendor(ctx.Request().Context(), id, user.CreatedBy)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, map[string]string{"message": "kontak vendor deleted"}, nil)
	}
}
