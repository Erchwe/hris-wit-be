package application

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/attachments/service"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

func AddRouteAttachment(s *httpservice.Service, cfg config.KVStore, e *echo.Echo) {
	svc := service.NewAttachmentService(s.GetDB(), cfg)

	attachment := e.Group("/attachment")

	attachment.GET("", getListAttachment(svc))
	attachment.GET("/:id", getAttachmentByID(svc))
	attachment.POST("", insertAttachment(svc, cfg))
	attachment.PUT("/:id", updateAttachment(svc, cfg))
	attachment.DELETE("/:id", deleteAttachment(svc))
}

func insertAttachment(svc *service.AttachmentService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request payload.InsertAttachmentPayload
		if err := ctx.Bind(&request); err != nil {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		err := svc.InsertAttachment(ctx.Request().Context(), request, user, cfg)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "attachment berhasil ditambahkan"}, nil)
	}
}

func getListAttachment(svc *service.AttachmentService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		result, err := svc.GetListAttachments(ctx.Request().Context())
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func getAttachmentByID(svc *service.AttachmentService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		result, err := svc.GetAttachmentByID(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, result, nil)
	}
}

func updateAttachment(svc *service.AttachmentService, cfg config.KVStore) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		var request payload.UpdateAttachmentPayload
		if err := ctx.Bind(&request); err != nil {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		if err := request.Validate(); err != nil {
			return err
		}

		user := sqlc.GetUserBackofficeRow{}
		err := svc.UpdateAttachment(ctx.Request().Context(), request, user)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "attachment berhasil diupdate"}, nil)
	}
}

func deleteAttachment(svc *service.AttachmentService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		if id == "" {
			return errors.WithStack(httpservice.ErrBadRequest)
		}

		err := svc.DeleteAttachment(ctx.Request().Context(), id)
		if err != nil {
			return err
		}

		return httpservice.ResponseData(ctx, map[string]string{"message": "attachment berhasil dihapus"}, nil)
	}
}
