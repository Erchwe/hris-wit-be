package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/data_pegawai/service" // Arahkan ke service pegawai yang baru
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
	"net/http"
)

func AddRouteDataPegawai(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewDataPegawaiService(s.GetDB(), cfg)

	pegawai := e.Group("/pegawai")

	pegawai.POST("", insertDataPegawai(svc))
	pegawai.GET("", getAllDataPegawai(svc))

}

func insertDataPegawai(svc *service.DataPegawaiService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.InsertDataPegawaiPayload

		// Bind request body
		if err := ctx.Bind(&request); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "failed to parse request")
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		// Validasi payload
		if err := request.Validate(); err != nil {
			log.FromCtx(ctx.Request().Context()).Error(err, "validation failed")
			return err
		}

		// Panggil service untuk insert
		_, err := svc.InsertDataPegawai(ctx.Request().Context(), request)
		if err != nil {
			return err // Kembalikan error dari service
		}

		// Kirim respons sukses
		return httpservice.ResponseData(ctx, map[string]string{"message": "Data pegawai berhasil ditambahkan"}, nil)
	}
}
func getAllDataPegawai(svc *service.DataPegawaiService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		pegawaiList, err := svc.GetAllDataPegawai(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"status":  "error",
				"message": "Gagal mengambil data pegawai",
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"status": "success",
			"data":   pegawaiList,
		})
	}
}
