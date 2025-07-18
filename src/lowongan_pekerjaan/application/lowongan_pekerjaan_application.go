package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/lowongan_pekerjaan/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
	"net/http"
)

func AddRouteLowonganPekerjaan(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewLowonganPekerjaanService(s.GetDB(), s.Cfg)

	lowongan := e.Group("/lowongan")

	//insert
	lowongan.POST("/insert", insertLowonganPekerjaan(svc, cfg))
	//delete
	lowongan.DELETE("/:id", deleteLowonganPekerjaan(svc))
	//list
	lowongan.GET("", listLowonganPekerjaan(svc))
	//GET
	lowongan.GET("/:id", GetLowonganPekerjaan(svc))
	//Update
	lowongan.PUT("/:id", updateLowonganPekerjaan(svc))

}

func insertLowonganPekerjaan(svc *service.LowonganPekerjaanService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.InsertLowonganPekerjaanPayload

		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request body")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}
		LowonganPekerjaan, err := svc.InsertLowonganPekerjaan(ctx.Request().Context(), request)
		if err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to insert lowongan pekerjaan")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}

		return httpservice.ResponseData(ctx, map[string]string{
			"message":               "Lowongan pekerjaan berhasil ditambahkan",
			"id_lowongan_pekerjaan": LowonganPekerjaan.IDLowonganPekerjaan,
		}, nil)
	}
}

func listLowonganPekerjaan(svc *service.LowonganPekerjaanService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.ListLowonganPekerjaan(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func GetLowonganPekerjaan(svc *service.LowonganPekerjaanService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idLowongan := ctx.Param("id")

		lowongan, err := svc.GetLowonganPekerjaan(ctx.Request().Context(), idLowongan)
		if err != nil {
			return errors.WithStack(httpservice.ErrUnknownSource)
		}

		return httpservice.ResponseData(ctx, lowongan, nil)
	}
}

func deleteLowonganPekerjaan(svc *service.LowonganPekerjaanService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.Wrap(httpservice.ErrBadRequest, httpservice.MsgInvalidIDParam)
		}

		//userBackoffice := ctx.Get(constants.MddwUserBackoffice).(sqlc.GetUserBackofficeByEmailRow)

		err := svc.DeleteLowonganPekerjaan(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "Lowongan pekerjaan berhasil dihapus"}, nil)
	}
}

func updateLowonganPekerjaan(svc *service.LowonganPekerjaanService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idLowongan := ctx.Param("id")
		if idLowongan == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
		}

		var request payload.UpdateLowonganPekerjaanPayload
		if err := ctx.Bind(&request); err != nil {
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		request.IDLowonganPekerjaan = idLowongan

		// Ambil data lama dari DB
		oldLowongan, err := svc.GetLowonganPekerjaan(ctx.Request().Context(), idLowongan)
		if err != nil {
			return err
		}

		// Buat param update dengan data lama + perubahan baru
		updateParams, err := request.ToEntity(oldLowongan)
		if err != nil {
			return err
		}

		updatedLowongan, err := svc.UpdateLowonganPekerjaan(ctx.Request().Context(), updateParams)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, updatedLowongan, nil)
	}
}
