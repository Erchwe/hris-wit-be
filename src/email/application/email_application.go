package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/email/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

func AddRouteEmail(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	emailSvc := service.NewEmailService(cfg)

	emailGroup := e.Group("/send-report-email")

	emailGroup.POST("", sendReportEmail(emailSvc, cfg))
}

func sendReportEmail(svc *service.EmailService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.SendReportEmailPayload
		if err := ctx.Bind(&request); err != nil {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		err := svc.SendReportEmail(ctx.Request().Context(), request)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "Email laporan berhasil dikirim"}, nil)
	}
}
