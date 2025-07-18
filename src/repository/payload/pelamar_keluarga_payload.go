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

type InsertPelamarKeluargaPayload struct {
	//IdKeluarga         string `json:"id_keluarga" valid:"required"`
	IdPelamar          string `json:"id_pelamar"`
	NamaIstriSuami     string `json:"nama_istri_suami"`
	JenisKelamin       string `json:"jenis_kelamin" valid:"required,in(Laki-laki|Perempuan)"`
	TempatLahir        string `json:"tempat_lahir"`
	TglLahir           string `json:"tgl_lahir"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	PekerjaanSkrg      string `json:"pekerjaan_skrg"`
	AlamatRumah        string `json:"alamat_rumah"`
}

func (payload *InsertPelamarKeluargaPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertPelamarKeluargaPayload) ToEntity(cfg config.KVStore) (data sqlc.CreatePelamarKeluargaParams, err error) {
	var (
		userId = utility.GenerateGoogleUUID()
	)

	data = sqlc.CreatePelamarKeluargaParams{
		IDKeluarga:         userId,
		IDPelamar:          payload.IdPelamar,
		NamaIstriSuami:     sql.NullString{},
		JenisKelamin:       payload.JenisKelamin,
		TempatLahir:        sql.NullString{},
		TglLahir:           sql.NullTime{},
		PendidikanTerakhir: sql.NullString{},
		PekerjaanSkrg:      sql.NullString{},
		AlamatRumah:        sql.NullString{},
		CreatedBy:          userId,
	}

	if payload.NamaIstriSuami != "" {
		data.NamaIstriSuami = sql.NullString{
			String: payload.NamaIstriSuami,
			Valid:  true,
		}
	}
	if payload.TempatLahir != "" {
		data.TempatLahir = sql.NullString{
			String: payload.TempatLahir,
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
	return
}
