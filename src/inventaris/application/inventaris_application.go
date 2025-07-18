package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/common/helper"
	"github.com/wit-id/blueprint-backend-go/src/inventaris/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
	"net/http"
)

func AddRouteInventaris(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewInventarisService(s.GetDB(), cfg)

	inv := e.Group("/inventaris")
	inv.POST("/insert", insertInventaris(svc))
	inv.GET("", getListInventaris(svc))
	inv.GET("/:id", getInventarisByID(svc))
	inv.GET("/with-relations", getListInventarisWithRelations(svc))
	inv.GET("/with-relations/:id", getInventarisWithRelationsByID(svc))
	inv.PUT("/update:id", updateInventaris(svc, cfg))
	inv.DELETE("/delete:id", deleteInventaris(svc))
}

func insertInventaris(svc *service.InventarisService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request payload.CreateInventarisPayload
		if err := c.Bind(&request); err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed to bind request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		// Validasi payload
		if err := request.Validate(); err != nil {
			return err
		}

		fileHeader, err := c.FormFile("image_file")
		if err == nil {
			savePath, err := helper.SaveFile(fileHeader, "inventaris")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyimpan gambar"})
			}
			baseURL := "http://localhost:6969"
			request.ImageURL = baseURL + savePath
		}

		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.CreateInventaris(c.Request().Context(), request, user)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(c, result, nil)
	}
}

func getListInventaris(svc *service.InventarisService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListInventaris(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListInventarisWithRelations(svc *service.InventarisService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListInventarisWithRelations(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getInventarisByID(svc *service.InventarisService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetInventarisByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getInventarisWithRelationsByID(svc *service.InventarisService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetInventarisWithRelationsByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateInventaris(svc *service.InventarisService, cfg config.KVStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request payload.UpdateInventarisPayload
		if err := c.Bind(&request); err != nil {
			log.FromCtx(c.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		fileHeader, err := c.FormFile("image_file")
		if err == nil {
			savePath, err := helper.SaveFile(fileHeader, "inventaris")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyimpan gambar"})
			}
			request.ImageURL = savePath
		}

		user := sqlc.GetUserBackofficeRow{}
		result, err := svc.UpdateInventaris(c.Request().Context(), request, user, cfg)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(c, result, nil)
	}
}

func deleteInventaris(svc *service.InventarisService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		user := sqlc.GetUserBackofficeRow{}
		err := svc.DeleteInventaris(ctx.Request().Context(), id, user)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "inventaris deleted"}, nil)
	}
}
