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

type InsertPegawaiPenguasaanBahasaPayload struct {
	//IdBahasa   string `json:"id_bahasa" valid:"required"`
	IdPegawai  string `json:"id_pegawai" valid:"required"`
	Bahasa     string `json:"bahasa" valid:"required"`
	Lisan      string `json:"lisan" valid:"required"`
	Tulisan    string `json:"tulisan" valid:"required"`
	Keterangan string `json:"keterangan"`
}

func (payload *InsertPegawaiPenguasaanBahasaPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return
}

func (payload *InsertPegawaiPenguasaanBahasaPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeByEmailRow) (data sqlc.CreatePegawaiPenguasaanBahasaParams) {
	data = sqlc.CreatePegawaiPenguasaanBahasaParams{
		IDBahasa:   utility.GenerateGoogleUUID(),
		IDPegawai:  payload.IdPegawai,
		Bahasa:     payload.Bahasa,
		Lisan:      sql.NullString{},
		Tulisan:    sql.NullString{},
		Keterangan: sql.NullString{},
		CreatedBy:  userData.CreatedBy,
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
