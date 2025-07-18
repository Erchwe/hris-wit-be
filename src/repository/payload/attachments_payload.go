package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type InsertAttachmentPayload struct {
	AttachID   string `json:"attach_id" valid:"required"`
	TaskID     string `json:"task_id" valid:"required"`
	AttachName string `json:"attach_name"`
	AttachURL  string `json:"attach_url"`
	CreatedBy  string `json:"created_by" valid:"required"`
}

func (payload *InsertAttachmentPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertAttachmentPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeRow) (data sqlc.CreateAttachmentParams) {
	data = sqlc.CreateAttachmentParams{
		AttachID:   payload.AttachID,
		TaskID:     payload.TaskID,
		AttachName: sql.NullString{}, // Nullable field
		AttachUrl:  sql.NullString{}, // Nullable field
		CreatedBy:  userData.CreatedBy,
	}

	if payload.AttachName != "" {
		data.AttachName = sql.NullString{
			String: payload.AttachName,
			Valid:  true,
		}
	}

	if payload.AttachURL != "" {
		data.AttachUrl = sql.NullString{
			String: payload.AttachURL,
			Valid:  true,
		}
	}

	return
}

type UpdateAttachmentPayload struct {
	AttachID   string `json:"attach_id" valid:"required"`
	TaskID     string `json:"task_id"`
	AttachName string `json:"attach_name"`
	AttachURL  string `json:"attach_url"`
	UpdatedBy  string `json:"updated_by" valid:"required"`
}

func (payload *UpdateAttachmentPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *UpdateAttachmentPayload) ToEntity(userData sqlc.GetUserBackofficeRow) (data sqlc.UpdateAttachmentParams) {
	data = sqlc.UpdateAttachmentParams{
		AttachID:   payload.AttachID,
		TaskID:     payload.TaskID,
		AttachName: sql.NullString{}, // Nullable field
		AttachUrl:  sql.NullString{}, // Nullable field
		UpdatedBy: sql.NullString{
			String: userData.UpdatedBy.String,
			Valid:  true,
		},
	}

	if payload.AttachName != "" {
		data.AttachName = sql.NullString{
			String: payload.AttachName,
			Valid:  true,
		}
	}

	if payload.AttachURL != "" {
		data.AttachUrl = sql.NullString{
			String: payload.AttachURL,
			Valid:  true,
		}
	}

	return
}
