package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/brand/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteBrand(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewBrandService(s.GetDB(), cfg)

	route := e.Group("/brand")
	route.POST("", insertBrand(svc))
	route.GET("", getListBrand(svc))
	route.GET("/:id", getBrandByID(svc))
	route.PUT("/:id", updateBrand(svc))
	route.DELETE("/:id", deleteBrand(svc))
}

func insertBrand(svc *service.BrandService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreateBrandPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.CreateBrand(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListBrand(svc *service.BrandService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListBrand(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getBrandByID(svc *service.BrandService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetBrandByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateBrand(svc *service.BrandService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.UpdateBrandPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}

		result, err := svc.UpdateBrand(ctx.Request().Context(), request)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deleteBrand(svc *service.BrandService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		user := sqlc.GetUserBackofficeRow{}
		err := svc.DeleteBrand(ctx.Request().Context(), id, user.CreatedBy)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, map[string]string{"message": "brand deleted"}, nil)
	}
}
