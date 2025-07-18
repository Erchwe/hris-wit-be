package payload

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type InsertPelamarPendidikanNonFormalPayload struct {
	//IdPddkNonFormal string  `json:"id_pddk_non_formal" valid:"required"`
	IdPelamar       string  `json:"id_pelamar"`
	Institusi       string  `json:"institusi"`
	JenisPendidikan string  `json:"jenis_pendidikan"`
	Kota            string  `json:"kota"`
	TglLulus        *string `json:"tgl_lulus"`
}

func (payload *InsertPelamarPendidikanNonFormalPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertPelamarPendidikanNonFormalPayload) ToEntity(cfg config.KVStore) (data sqlc.CreatePelamarPendidikanNonFormalParams, err error) {
	var (
		userId = utility.GenerateGoogleUUID()
	)

	data = sqlc.CreatePelamarPendidikanNonFormalParams{
		IDPddkNonFormal: userId,
		IDPelamar:       payload.IdPelamar,
		Institusi:       payload.Institusi,
		JenisPendidikan: payload.JenisPendidikan,
		Kota:            payload.Kota,
		TglLulus:        sql.NullTime{},
		CreatedBy:       userId,
	}

	if payload.TglLulus != nil {
		if parsed, errParse := time.Parse("2006-01-02", *payload.TglLulus); errParse == nil {
			data.TglLulus = sql.NullTime{Time: parsed, Valid: true}
		} else {
			err = errors.Wrapf(httpservice.ErrBadRequest, "invalid tgl_lulus format: %s", errParse.Error())
			return
		}
	}

	return
}
