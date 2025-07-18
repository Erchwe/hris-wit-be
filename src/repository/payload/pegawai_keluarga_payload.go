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

type InsertPegawaiKeluargaPayload struct {
	//IdKeluarga         string `json:"id_keluarga" valid:"required"`
	IdPegawai          string `json:"id_pegawai" valid:"required"`
	NamaIstriSuami     string `json:"nama_istri_suami" valid:"required"`
	JenisKelamin       string `json:"jenis_kelamin" valid:"required"`
	TempatLahir        string `json:"tempat_lahir"`
	TglLahir           string `json:"tgl_lahir"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	PekerjaanSkrg      string `json:"pekerjaan_skrg"`
	AlamatRumah        string `json:"alamat_rumah"`
}

func (payload *InsertPegawaiKeluargaPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return
}

func (payload *InsertPegawaiKeluargaPayload) ToEntity(cfg config.KVStore, userData sqlc.GetUserBackofficeByEmailRow) (data sqlc.CreatePegawaiKeluargaParams, err error) {
	data = sqlc.CreatePegawaiKeluargaParams{
		IDKeluarga:         utility.GenerateGoogleUUID(),
		IDPegawai:          payload.IdPegawai,
		NamaIstriSuami:     sql.NullString{},
		JenisKelamin:       nil,
		TempatLahir:        sql.NullString{},
		TglLahir:           sql.NullTime{},
		PendidikanTerakhir: sql.NullString{},
		PekerjaanSkrg:      sql.NullString{},
		AlamatRumah:        sql.NullString{},
		CreatedBy:          userData.CreatedBy,
	}

	if payload.TempatLahir != "" {
		data.TempatLahir = sql.NullString{
			String: payload.TempatLahir,
			Valid:  true,
		}
	}
	if payload.PendidikanTerakhir != "" {
		data.PendidikanTerakhir = sql.NullString{
			String: payload.PendidikanTerakhir,
			Valid:  true,
		}
	}
	if payload.PekerjaanSkrg != "" {
		data.PekerjaanSkrg = sql.NullString{
			String: payload.PekerjaanSkrg,
			Valid:  true,
		}
	}
	if payload.AlamatRumah != "" {
		data.AlamatRumah = sql.NullString{
			String: payload.AlamatRumah,
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

	return
}
