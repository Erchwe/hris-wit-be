package payload

import (
	_ "database/sql"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type InsertProjectMemberPayload struct {
	ProjectMembID string `json:"project_memb_id" valid:"required"`
	IDDataPegawai string `json:"id_data_pegawai" valid:"required"`
	ProjectID     string `json:"project_id" valid:"required"`
	ProjectRole   string `json:"project_role" valid:"required"`
	CreatedBy     string `json:"created_by" valid:"required"`
}

func (payload *InsertProjectMemberPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertProjectMemberPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeRow) (data sqlc.CreateProjectMemberParams) {
	data = sqlc.CreateProjectMemberParams{
		ProjectMembID: payload.ProjectMembID,
		IDDataPegawai: payload.IDDataPegawai,
		ProjectID:     payload.ProjectID,
		ProjectRole:   payload.ProjectRole,
	}

	return
}
