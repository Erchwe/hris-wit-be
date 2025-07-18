package payload

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type InsertPelamarPenguasaanBahasaPayload struct {
	//IdBahasa   string `json:"id_bahasa" valid:"required"`
	IdPelamar  string `json:"id_pelamar"`
	Bahasa     string `json:"bahasa"`
	Lisan      string `json:"lisan"`
	Tulisan    string `json:"tulisan"`
	Keterangan string `json:"keterangan"`
}

func (payload *InsertPelamarPenguasaanBahasaPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertPelamarPenguasaanBahasaPayload) ToEntity(cfg config.KVStore) (data sqlc.CreatePelamarPenguasaanBahasaParams) {
	var (
		userId = utility.GenerateGoogleUUID()
	)

	data = sqlc.CreatePelamarPenguasaanBahasaParams{
		IDBahasa:   userId,
		IDPelamar:  payload.IdPelamar,
		Bahasa:     payload.Bahasa,
		Lisan:      sql.NullString{},
		Tulisan:    sql.NullString{},
		Keterangan: sql.NullString{},
		CreatedBy:  userId,
	}

	if payload.Lisan != "" {
		data.Lisan = sql.NullString{
			String: payload.Lisan,
			Valid:  true,
		}
	}

	if payload.Tulisan != "" {
		data.Tulisan = sql.NullString{
			String: payload.Tulisan,
			Valid:  true,
		}
	}

	if payload.Keterangan != "" {
		data.Keterangan = sql.NullString{
			String: payload.Keterangan,
			Valid:  true,
		}
	}

	return
}
