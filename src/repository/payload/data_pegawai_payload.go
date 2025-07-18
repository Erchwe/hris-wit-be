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

type InsertDataPegawaiPayload struct {
	IdDataPegawai   string `json:"id_data_pegawai" valid:"required"`
	EmployeeNumber  string `json:"employee_number"`
	Divisi          string `json:"divisi"`
	NamaLengkap     string `json:"nama_lengkap"`
	TempatLahir     string `json:"tempat_lahir"`
	TglLahir        string `json:"tgl_lahir"`
	JenisKelamin    string `json:"jenis_kelamin"`
	Kewarganegaraan string `json:"kewarganegaraan"`
	Phone           string `json:"phone"`
	Mobile          string `json:"mobile"`
	Agama           string `json:"agama"`
	GolDarah        string `json:"gol_darah"`
	Gaji            string `json:"gaji"`
	StatusMenikah   *bool  `json:"status_menikah"`
	NoKTP           string `json:"no_ktp"`
	NoNPWP          string `json:"no_npwp"`
	Status          string `json:"status"`
}

func (payload *InsertDataPegawaiPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertDataPegawaiPayload) ToEntity(cfg config.KVStore) (data sqlc.CreateDataPegawaiParams) {
	tglLahir, err := time.Parse("2006-01-02", payload.TglLahir)
	if err != nil {
		return data
	}

	data = sqlc.CreateDataPegawaiParams{
		IDDataPegawai:   utility.GenerateGoogleUUID(),
		EmployeeNumber:  payload.EmployeeNumber,
		Divisi:          sql.NullString{String: payload.Divisi, Valid: payload.Divisi != ""},
		NamaLengkap:     payload.NamaLengkap,
		TempatLahir:     sql.NullString{String: payload.TempatLahir, Valid: payload.TempatLahir != ""},
		TglLahir:        sql.NullTime{Time: tglLahir, Valid: true},
		JenisKelamin:    payload.JenisKelamin,
		Kewarganegaraan: sql.NullString{String: payload.Kewarganegaraan, Valid: payload.Kewarganegaraan != ""},
		Phone:           sql.NullString{String: payload.Phone, Valid: payload.Phone != ""},
		Mobile:          sql.NullString{String: payload.Mobile, Valid: payload.Mobile != ""},
		Agama:           sql.NullString{String: payload.Agama, Valid: payload.Agama != ""},
		GolDarah:        sql.NullString{String: payload.GolDarah, Valid: payload.GolDarah != ""},
		Gaji:            sql.NullString{String: payload.Gaji, Valid: payload.Gaji != ""},
		NoKtp:           sql.NullString{String: payload.NoKTP, Valid: payload.NoKTP != ""},
		NoNpwp:          sql.NullString{String: payload.NoNPWP, Valid: payload.NoNPWP != ""},
		Status:          sql.NullString{String: payload.Status, Valid: payload.Status != ""},
		CreatedBy:       payload.IdDataPegawai,
	}

	// Optional fields
	if payload.Phone != "" {
		data.Phone = sql.NullString{
			String: payload.Phone,
			Valid:  true,
		}
	}

	if payload.Mobile != "" {
		data.Mobile = sql.NullString{
			String: payload.Mobile,
			Valid:  true,
		}
	}
	if payload.Agama != "" {
		data.Agama = sql.NullString{
			String: payload.Agama,
			Valid:  true,
		}
	}
	if payload.GolDarah != "" {
		data.GolDarah = sql.NullString{
			String: payload.GolDarah,
			Valid:  true,
		}
	}

	if payload.Gaji != "" {
		data.Gaji = sql.NullString{
			String: payload.Gaji,
			Valid:  true,
		}
	}

	if payload.StatusMenikah != nil {
		data.StatusMenikah = sql.NullBool{
			Bool:  *payload.StatusMenikah,
			Valid: true,
		}
	}

	if payload.NoKTP != "" {
		data.NoKtp = sql.NullString{
			String: payload.NoKTP,
			Valid:  true,
		}
	}

	if payload.NoNPWP != "" {
		data.NoNpwp = sql.NullString{
			String: payload.NoNPWP,
			Valid:  true,
		}
	}

	return
}
