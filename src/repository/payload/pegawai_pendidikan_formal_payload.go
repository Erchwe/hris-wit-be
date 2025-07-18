package payload

import (
	"database/sql"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type InsertPegawaiPendidikanFormalPayload struct {
	//IdPddkFormal    string   `json:"id_pddk_formal" valid:"required"`
	IdPegawai       string   `json:"id_pegawai" valid:"required"`
	JenjangPddk     string   `json:"jenjang_pddk" valid:"required"`
	NamaSekolah     string   `json:"nama_sekolah" valid:"required"`
	JurusanFakultas string   `json:"jurusan_fakultas"`
	Kota            string   `json:"kota"`
	TglLulus        *string  `json:"tgl_lulus"`
	IPK             *float64 `json:"ipk"`
}

func (payload *InsertPegawaiPendidikanFormalPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertPegawaiPendidikanFormalPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeByEmailRow) (data sqlc.CreatePegawaiPendidikanFormalParams, err error) {
	data = sqlc.CreatePegawaiPendidikanFormalParams{
		IDPddkFormal:    utility.GenerateGoogleUUID(),
		IDPegawai:       payload.IdPegawai,
		JenjangPddk:     payload.JenjangPddk,
		NamaSekolah:     payload.NamaSekolah,
		JurusanFakultas: sql.NullString{},
		Kota:            sql.NullString{},
		TglLulus:        sql.NullTime{},
		Ipk:             sql.NullFloat64{},
		CreatedBy:       userData.CreatedBy,
	}

	if payload.JurusanFakultas != "" {
		data.JurusanFakultas = sql.NullString{
			String: payload.JurusanFakultas,
			Valid:  true,
		}
	}

	if payload.Kota != "" {
		data.Kota = sql.NullString{
			String: payload.Kota,
			Valid:  true,
		}
	}

	if payload.TglLulus != nil {
		if parsed, errParse := time.Parse("2006-01-02", *payload.TglLulus); errParse == nil {
			data.TglLulus = sql.NullTime{Time: parsed, Valid: true}
		} else {
			err = errors.Wrapf(httpservice.ErrBadRequest, "invalid tgl_lulus format: %s", errParse.Error())
			return
		}
	}

	if payload.IPK != nil {
		data.Ipk = sql.NullFloat64{
			Float64: *payload.IPK,
			Valid:   true,
		}
	}

	return
}
