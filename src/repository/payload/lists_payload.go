package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type InsertListPayload struct {
	ProjectID string `json:"project_id" valid:"required"`
	ListName  string `json:"list_name" valid:"required"`
	ListOrder int64  `json:"list_order"`
	CreatedBy string `json:"created_by" valid:"required"`
}

func (payload *InsertListPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertListPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeRow) (data sqlc.CreateListParams) {
	data = sqlc.CreateListParams{
		ProjectID: payload.ProjectID,
		ListName:  sql.NullString{},
		ListOrder: sql.NullInt64{Int64: payload.ListOrder, Valid: true},
		CreatedBy: userData.CreatedBy,
	}

	if payload.ListName != "" {
		data.ListName = sql.NullString{
			String: payload.ListName,
			Valid:  true,
		}
	}
	return
}

type UpdateListPayload struct {
	ListID    string `json:"list_id" valid:"required"`
	ListName  string `json:"list_name"`
	ListOrder int64  `json:"list_order"`
	UpdatedBy string `json:"updated_by" valid:"required"`
}

func (payload *UpdateListPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *UpdateListPayload) ToEntity(userData sqlc.GetUserBackofficeRow) (data sqlc.UpdateListParams) {
	data = sqlc.UpdateListParams{
		ListID:    payload.ListID,
		ListName:  sql.NullString{String: payload.ListName, Valid: payload.ListName != ""},
		ListOrder: sql.NullInt64{Int64: payload.ListOrder, Valid: true}, // Set ListOrder
		UpdatedBy: sql.NullString{
			String: userData.UpdatedBy.String,
			Valid:  true,
		},
	}
	if payload.ListName != "" {
		data.ListName = sql.NullString{
			String: payload.ListName,
			Valid:  true,
		}
	}
	return
}
