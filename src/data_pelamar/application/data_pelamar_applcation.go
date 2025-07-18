package application

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/data_pelamar/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func AddRouteUserDataPelamar(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewDataPelamarService(s.GetDB(), cfg)

	pelamar := e.Group("/pelamar")

	// GET List Pelamar
	pelamar.GET("", listPelamar(svc))
	// GET Pelamar id
	pelamar.GET("/:id", getPelamarByID(svc))
	// POST Insert Pelamar
	pelamar.POST("/insert", insertPelamarLengkap(svc, cfg))
	// PUT Update Pelamar
	pelamar.PUT("/update-status", updateStatusPelamar(svc))
}

// insertPelamarLengkap
func insertPelamarLengkap(svc *service.DataPelamarService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.InsertPelamarLengkapPayload

		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "validation failed")
			return err
		}

		err := svc.InsertDataPelamarLengkap(ctx.Request().Context(), request)
		if err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to insert data pelamar")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
		fmt.Println("d")

		return httpservice.ResponseData(ctx, map[string]string{"message": "Data pelamar berhasil ditambahkan"}, nil)
	}
}

// getListPelamar
func listPelamar(svc *service.DataPelamarService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.ListPelamar(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

// getPelamarByID handles retrieving a specific pelamar by ID.
func getPelamarByID(svc *service.DataPelamarService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetPelamarByID(ctx.Request().Context(), id)
		if err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to get pelamar by ID")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getPelamarByEmail(svc *service.DataPelamarService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		email := ctx.Param("email")
		if email == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetPelamarByEmail(ctx.Request().Context(), email)
		if err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to get pelamar by email")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

// Handler baru untuk batch update status
func updateStatusPelamar(svc *service.DataPelamarService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		type requestPayload struct {
			IDs    []string `json:"ids"`
			Status string   `json:"status"` // Menerima status dalam format Base64
		}

		var req requestPayload
		if err := ctx.Bind(&req); err != nil {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		// ✅ FIX: Lakukan decode Base64 di sini
		decodedStatus, err := base64.StdEncoding.DecodeString(req.Status)
		if err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "gagal decode base64 status")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		statusString := string(decodedStatus)

		for _, id := range req.IDs {
			params := sqlc.UpdatePelamarParams{
				IDDataPelamar: id,
				Status: sql.NullString{
					String: statusString,
					Valid:  true,
				},
			}

			_, err := svc.UpdateStatusPelamar(ctx.Request().Context(), params)
			if err != nil {
				log.FromCtx(ctx.Request().Context()).Error(err, "gagal update status untuk id", "pelamar_id", id)
				return errors.WithStack(httpservice.ErrUnknownSource)
			}
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "Status berhasil diupdate"}, nil)
	}
}
