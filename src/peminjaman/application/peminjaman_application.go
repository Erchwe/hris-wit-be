package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/peminjaman/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

// AddRoutePeminjaman: Ini adalah versi dari Anda, saya tidak akan mengubahnya drastis
// kecuali untuk memastikan updatePeminjaman dipanggil dengan benar.
func AddRoutePeminjaman(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewPeminjamanService(s.GetDB(), cfg)

	p := e.Group("/peminjaman")
	p.POST("", insertPeminjaman(svc))
	p.GET("", getListPeminjaman(svc))
	p.GET("/:id", getPeminjamanByID(svc))
	p.PUT("/:id", updatePeminjaman(svc)) // Sudah benar
	p.DELETE("/:id", deletePeminjaman(svc))
}

// Helper untuk mendapatkan user dari context
func getUserFromContext(c echo.Context) (sqlc.GetUserBackofficeRow, error) {
	user, ok := c.Get("user").(sqlc.GetUserBackofficeRow) // Pastikan kunci "user" dan tipe GetUserBackofficeRow sesuai
	if !ok {
		log.FromCtx(c.Request().Context()).Error(nil, "User not found in context or type mismatch")
		return sqlc.GetUserBackofficeRow{}, errors.WithStack(httpservice.ErrUserNotFound)
	}
	return user, nil
}

func insertPeminjaman(svc *service.PeminjamanService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.CreatePeminjamanPayload
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}

		user, err := getUserFromContext(ctx) // Ambil user dari context
		if err != nil {
			return err
		}

		result, err := svc.CreatePeminjaman(ctx.Request().Context(), request, user) // Teruskan user
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getListPeminjaman(svc *service.PeminjamanService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListPeminjaman(ctx.Request().Context())
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getPeminjamanByID(svc *service.PeminjamanService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetPeminjamanByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updatePeminjaman(svc *service.PeminjamanService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		peminjamanID := ctx.Param("id") // Ambil ID dari URL
		if peminjamanID == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		var request payload.UpdatePeminjamanPayload // <<--- GUNAKAN PAYLOAD BARU
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to bind update payload")
			return errors.WithStack(httpservice.ErrBadRequest)
		}
		if err := request.Validate(); err != nil {
			return err
		}

		user, err := getUserFromContext(ctx) // Ambil user dari context
		if err != nil {
			return err
		}

		result, err := svc.UpdatePeminjaman(ctx.Request().Context(), peminjamanID, request, user) // Teruskan ID dan user
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, result, nil)
	}
}

func deletePeminjaman(svc *service.PeminjamanService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		user, err := getUserFromContext(ctx) // Ambil user dari context
		if err != nil {
			return err
		}

		err = svc.DeletePeminjaman(ctx.Request().Context(), id, user) // Teruskan user
		if err != nil {
			return err
		}
		return httpservice.ResponseData(ctx, map[string]string{"message": "peminjaman deleted"}, nil)
	}
}
