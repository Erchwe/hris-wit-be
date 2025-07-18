package payload

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type InsertTaskMemberPayload struct {
	TaskMembID    string `json:"task_memb_id" valid:"required"`
	ProjectMembID string `json:"project_memb_id" valid:"required"`
	TaskID        string `json:"task_id" valid:"required"`
	CreatedBy     string `json:"created_by" valid:"required"`
}

func (payload *InsertTaskMemberPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertTaskMemberPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeRow) (data sqlc.CreateTaskMemberParams) {
	data = sqlc.CreateTaskMemberParams{
		TaskMembID:    payload.TaskMembID,
		ProjectMembID: payload.ProjectMembID,
		TaskID:        payload.TaskID,
	}

	return
}
