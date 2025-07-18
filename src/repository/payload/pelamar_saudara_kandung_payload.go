package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"time"
)

type InsertPelamarSaudaraKandungPayload struct {
	//IdSaudara           string `json:"id_saudara" valid:"required"`
	IdPelamar           string `json:"id_pelamar"`
	Nama                string `json:"nama_saudara_kandung" valid:"required"`
	JenisKelamin        string `json:"jenis_kelamin" valid:"required,in(Laki-laki|Perempuan)"`
	TempatLahir         string `json:"tempat_lahir"`
	PendidikanPekerjaan string `json:"pendidikan_pekerjaan"`
	TglLahir            string `json:"tgl_lahir"`
}

func (payload *InsertPelamarSaudaraKandungPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return nil
}

func (payload *InsertPelamarSaudaraKandungPayload) ToEntity(cfg config.KVStore) (data sqlc.CreatePelamarSaudaraKandungParams, err error) {
	var (
		userId = utility.GenerateGoogleUUID()
	)

	data = sqlc.CreatePelamarSaudaraKandungParams{
		IDSaudara:           userId,
		IDPelamar:           payload.IdPelamar,
		Nama:                payload.Nama,
		JenisKelamin:        payload.JenisKelamin,
		TempatLahir:         sql.NullString{},
		PendidikanPekerjaan: sql.NullString{},
		TglLahir:            sql.NullTime{},
		CreatedBy:           userId,
	}

	if payload.TempatLahir != "" {
		data.TempatLahir = sql.NullString{
			String: payload.TempatLahir,
			Valid:  true,
		}
	}

	if payload.PendidikanPekerjaan != "" {
		data.PendidikanPekerjaan = sql.NullString{
			String: payload.PendidikanPekerjaan,
			Valid:  true,
		}
	}

	if payload.TglLahir != "" {
		parsedTime, parseErr := time.Parse("2006-01-02", payload.TglLahir)
		if parseErr == nil {
			data.TglLahir = sql.NullTime{
				Time:  parsedTime,
				Valid: true,
			}
		}
	}

	return data, nil
}
