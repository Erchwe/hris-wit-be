package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/kategori/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteKategori(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewKategoriService(s.GetDB(), cfg)

	k := e.Group("/kategori")
	k.POST("/insert", insertKategori(svc))
	k.GET("", getListKategori(svc))
	k.GET("/:id", getKategoriByID(svc))
	k.PUT("/update:id", updateKategori(svc))
	k.DELETE("/delete:id", deleteKategori(svc))
}

func insertKategori(svc *service.KategoriService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreateKategoriPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.CreateKategori(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListKategori(svc *service.KategoriService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListKategori(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getKategoriByID(svc *service.KategoriService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		result, err := svc.GetKategoriByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateKategori(svc *service.KategoriService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.UpdateKategoriPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		result, err := svc.UpdateKategori(ctx.Request().Context(), request)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteKategori(svc *service.KategoriService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		user := sqlc.GetUserBackofficeRow{}
		if err := svc.DeleteKategori(ctx.Request().Context(), id, user.CreatedBy); err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, map[string]string{"message": "kategori deleted"}, nil)
	}
}
