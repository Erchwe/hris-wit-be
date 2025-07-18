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

type InsertPelamarReferensiPayload struct {
	//IdReferensi    string `json:"id_referensi" valid:"required"`
	IdPelamar      string `json:"id_pelamar"`
	Nama           string `json:"nama"`
	Namaperusahaan string `json:"nama_perusahaan"`
	Jabatan        string `json:"jabatan"`
	NoTelp         string `json:"no_telp_perusahaan"`
}

func (payload *InsertPelamarReferensiPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}

	return
}

func (payload *InsertPelamarReferensiPayload) ToEntity(cfg config.KVStore) (data sqlc.CreatePelamarReferensiParams) {
	var (
		userId = utility.GenerateGoogleUUID()
	)

	data = sqlc.CreatePelamarReferensiParams{
		IDReferensi:      userId,
		IDPelamar:        payload.IdPelamar,
		Nama:             payload.Nama,
		NamaPerusahaan:   sql.NullString{},
		Jabatan:          sql.NullString{},
		NoTelpPerusahaan: sql.NullString{},
		CreatedBy:        userId,
	}

	if payload.Namaperusahaan != "" {
		data.NamaPerusahaan = sql.NullString{
			String: payload.Namaperusahaan,
			Valid:  true,
		}
	}

	if payload.Jabatan != "" {
		data.Jabatan = sql.NullString{
			String: payload.Jabatan,
			Valid:  true,
		}
	}

	if payload.NoTelp != "" {
		data.NoTelpPerusahaan = sql.NullString{
			String: payload.NoTelp,
			Valid:  true,
		}
	}
	return
}
