package application

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/src/vendor_and_kontak/service"
	vendorservice "github.com/wit-id/blueprint-backend-go/src/vendors/service"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteVendorAndKontak(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewVendorAndKontakService(s.GetDB(), cfg, vendorservice.NewVendorService(s.GetDB(), cfg))
	grp := e.Group("/vendor-kontak")

	grp.POST("/insert", insertVendorWithKontak(svc))
	grp.PUT("/update/:id", updateVendorWithKontak(svc))
	grp.DELETE("/delete/:id", deleteVendorWithKontak(svc))
}

func insertVendorWithKontak(svc *service.VendorAndKontakService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreateVendorWithKontakPayload

		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{CreatedBy: "admin"} // ⛔ hardcoded
		vendor, kontakList, err := svc.CreateVendorWithKontak(ctx.Request().Context(), request, user, svc.Cfg)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]interface{}{
			"vendor": vendor,
			"kontak": kontakList,
		}, nil)
	}
}

func updateVendorWithKontak(svc *service.VendorAndKontakService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.UpdateVendorWithKontakPayload
		id := ctx.Param("id")

		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		request.Vendor.VendorID = id
		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{
			UpdatedBy: sql.NullString{String: "admin", Valid: true},
		}
		vendor, kontakList, err := svc.UpdateVendorWithKontak(ctx.Request().Context(), request, user, svc.Cfg)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]interface{}{
			"vendor": vendor,
			"kontak": kontakList,
		}, nil)
	}
}

func deleteVendorWithKontak(svc *service.VendorAndKontakService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		user := sqlc.GetUserBackofficeRow{
			DeletedBy: sql.NullString{String: "admin", Valid: true},
		}
		err := svc.DeleteVendorWithKontak(ctx.Request().Context(), id, user)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{
			"message": "Vendor dan kontak berhasil dihapus",
		}, nil)
	}
}
