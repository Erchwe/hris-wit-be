package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *AttachmentService) GetAttachmentByID(ctx context.Context, attachID string) (sqlc.Attachment, error) {
	q := sqlc.New(s.mainDB)

	attachment, err := q.GetAttachmentByID(ctx, attachID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get attachment by ID")
		return sqlc.Attachment{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return attachment, nil
}

func (s *AttachmentService) GetListAttachments(ctx context.Context) ([]sqlc.Attachment, error) {
	q := sqlc.New(s.mainDB)

	attachments, err := q.ListAttachments(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list of attachments")
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return attachments, nil
}
