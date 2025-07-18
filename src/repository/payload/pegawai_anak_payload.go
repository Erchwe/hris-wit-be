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

type InsertPegawaiAnakPayload struct {
	//IdAnak              string `json:"id_anak" valid:"required"`
	IdPegawai           string `json:"id_pegawai" valid:"required"`
	Nama                string `json:"nama" valid:"required"`
	JenisKelamin        string `json:"jenis_kelamin" valid:"required"`
	TempatLahir         string `json:"tempat_lahir"`
	TglLahir            string `json:"tgl_lahir"`
	PendidikanPekerjaan string `json:"pendidikan_pekerjaan"`
}

func (payload *InsertPegawaiAnakPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return
}

func (payload *InsertPegawaiAnakPayload) ToEntity(cfg config.KVStore) (data sqlc.CreatePegawaiAnakParams, err error) {
	var (
		userId = utility.GenerateGoogleUUID()
	)

	data = sqlc.CreatePegawaiAnakParams{
		IDAnak:              userId,
		IDPegawai:           payload.IdPegawai,
		Nama:                payload.Nama,
		JenisKelamin:        nil,
		TempatLahir:         sql.NullString{},
		TglLahir:            sql.NullTime{},
		PendidikanPekerjaan: sql.NullString{},
		CreatedBy:           userId,
	}

	if payload.TempatLahir != "" {
		data.TempatLahir = sql.NullString{
			String: payload.TempatLahir,
			Valid:  true,
		}
	}

	if payload.TglLahir != "" {
		if tgl, parseErr := time.Parse("2006-01-02", payload.TglLahir); parseErr == nil {
			data.TglLahir = sql.NullTime{Time: tgl, Valid: true}
		} else {
			err = errors.Wrapf(httpservice.ErrBadRequest, "invalid tgl_lahir format: %s", parseErr.Error())
			return
		}
	}

	if payload.PendidikanPekerjaan != "" {
		data.PendidikanPekerjaan = sql.NullString{
			String: payload.PendidikanPekerjaan,
			Valid:  true,
		}
	}

	return
}
