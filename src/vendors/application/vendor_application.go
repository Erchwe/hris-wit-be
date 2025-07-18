package application

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/src/vendors/service"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteVendor(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewVendorService(s.GetDB(), cfg)

	v := e.Group("/vendor")
	v.POST("/insert", insertVendor(svc))
	v.GET("", getListVendor(svc))
	v.GET("/:id", getVendorByID(svc))
	v.PUT("/update:id", updateVendor(svc))
	v.DELETE("/delete:id", deleteVendor(svc))
}

func insertVendor(svc *service.VendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreateVendorPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		fmt.Println("a")
		if err := request.Validate(); err != nil {
			return err
		}
		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.CreateVendor(ctx.Request().Context(), request, user)
		fmt.Println(result)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListVendor(svc *service.VendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		log.FromCtx(ctx.Request().Context()).Info("GET /vendor route triggered")
		result, err := svc.GetListVendors(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getVendorByID(svc *service.VendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		result, err := svc.GetVendorByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateVendor(svc *service.VendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.UpdateVendorPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		result, err := svc.UpdateVendor(ctx.Request().Context(), request)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteVendor(svc *service.VendorService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		user := sqlc.GetUserBackofficeRow{}
		if err := svc.DeleteVendor(ctx.Request().Context(), id, user.CreatedBy); err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, map[string]string{"message": "vendor deleted"}, nil)
	}
}
