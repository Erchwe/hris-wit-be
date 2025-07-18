package payload

import (
	"database/sql"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type InsertPegawaiPendidikanNonFormalPayload struct {
	//IdPddkNonFormal string `json:"id_pddk_non_formal" valid:"required"`
	IdPegawai       string `json:"id_pegawai" valid:"required"`
	Institusi       string `json:"institusi" valid:"required"`
	JenisPendidikan string `json:"jenis_pendidikan" valid:"required"`
	Kota            string `json:"kota"`
}

func (payload *InsertPegawaiPendidikanNonFormalPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return
}

func (payload *InsertPegawaiPendidikanNonFormalPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeByEmailRow) (data sqlc.CreatePegawaiPendidikanNonFormalParams, err error) {
	data = sqlc.CreatePegawaiPendidikanNonFormalParams{
		IDPddkNonFormal: utility.GenerateGoogleUUID(),
		IDPegawai:       payload.IdPegawai,
		Institusi:       payload.Institusi,
		JenisPendidikan: payload.JenisPendidikan,
		Kota:            sql.NullString{},
		CreatedBy:       userData.CreatedBy,
	}

	if payload.Kota != "" {
		data.Kota = sql.NullString{
			String: payload.Kota,
			Valid:  true,
		}
	}

	return
}
