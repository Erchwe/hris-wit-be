package payload

import (
	"database/sql"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"time"
)

type InsertDataPelamarPayload struct {
	IDDataPelamar       string  `json:"id_data_pelamar"`
	IdLowonganPekerjaan string  `json:"id_lowongan_pekerjaan" valid:"required"`
	Email               string  `json:"email" valid:"email,required"`
	NamaLengkap         string  `json:"nama_lengkap" valid:"required"`
	TempatLahir         string  `json:"tempat_lahir"`
	TglLahir            string  `json:"tgl_lahir"`
	JenisKelamin        string  `json:"jenis_kelamin" valid:"required,in(Laki-laki|Perempuan)"`
	Kewarganegaraan     string  `json:"kewarganegaraan"`
	Phone               string  `json:"phone"`
	Mobile              string  `json:"mobile"`
	Agama               string  `json:"agama"`
	GolDarah            string  `json:"gol_darah"`
	StatusMenikah       bool    `json:"status_menikah"`
	NoKTP               string  `json:"no_ktp" valid:"length(16|16)"`
	NoNPWP              string  `json:"no_npwp"`
	Status              string  `json:"status" valid:"in(New|Short list|HR Interview|User Interview|Refference Checking|Offering|Psikotest|Hired|Rejected)"`
	AsalKota            string  `json:"asal_kota"`
	GajiTerakhir        float64 `json:"gaji_terakhir"`
	HarapanGaji         float64 `json:"harapan_gaji"`
	SedangBekerja       string  `json:"sedang_bekerja"`
	KetersediaanBekerja string  `json:"ketersediaan_bekerja"`
	SumberInformasi     string  `json:"sumber_informasi"`
	Alasan              string  `json:"alasan"`
	KetersediaanInter   string  `json:"ketersediaan_inter"`
	ProfesiKerja        string  `json:"profesi_kerja" valid:"in(Junior|Senior|Manager)"`
}

func (payload *InsertDataPelamarPayload) Validate() (err error) {
	if payload.IdLowonganPekerjaan == "" {
		return errors.Wrap(httpservice.ErrBadRequest, "id_lowongan_pekerjaan is required")
	}

	if payload.Email == "" {
		return errors.Wrap(httpservice.ErrBadRequest, "email is required")
	}

	if !govalidator.IsEmail(payload.Email) {
		return errors.Wrap(httpservice.ErrBadRequest, "invalid email format")
	}

	if payload.NamaLengkap == "" {
		return errors.Wrap(httpservice.ErrBadRequest, "nama_lengkap is required")
	}

	if payload.JenisKelamin == "" {
		return errors.Wrap(httpservice.ErrBadRequest, "jenis_kelamin is required")
	}

	return nil
}

func (payload *InsertDataPelamarPayload) ToEntity(cfg config.KVStore) (data sqlc.CreatePelamarParams) {

	data = sqlc.CreatePelamarParams{
		IDDataPelamar:       payload.IDDataPelamar,
		IDLowonganPekerjaan: payload.IdLowonganPekerjaan,
		Email:               payload.Email,
		NamaLengkap:         payload.NamaLengkap,
		TempatLahir:         sql.NullString{}, // ini data kosong
		TglLahir:            sql.NullTime{},   // ini data kosong atau sama jaa kayak gaada
		JenisKelamin:        payload.JenisKelamin,
		Kewarganegaraan:     sql.NullString{},
		Phone:               sql.NullString{},
		Mobile:              sql.NullString{},
		Agama:               sql.NullString{},
		GolDarah:            sql.NullString{},
		StatusMenikah:       sql.NullBool{},
		NoKtp:               sql.NullString{},
		NoNpwp:              sql.NullString{},
		Status:              payload.Status,
		AsalKota:            sql.NullString{},
		GajiTerakhir:        sql.NullString{},
		HarapanGaji:         sql.NullString{},
		SedangBekerja:       sql.NullString{},
		KetersediaanBekerja: sql.NullTime{},
		SumberInformasi:     sql.NullString{},
		Alasan:              sql.NullString{},
		KetersediaanInter:   sql.NullTime{},
		ProfesiKerja:        payload.ProfesiKerja,
		CreatedBy:           payload.IDDataPelamar,
	}
	fmt.Println("b")

	if payload.TempatLahir != "" {
		data.TempatLahir = sql.NullString{
			String: payload.TempatLahir,
			Valid:  true,
		}
	}
	if payload.TglLahir != "" {
		parsedTime, err := time.Parse("2006-01-02", payload.TglLahir)
		if err == nil {
			data.TglLahir = sql.NullTime{
				Time: parsedTime, Valid: true,
			}
		}
	}
	if payload.Kewarganegaraan != "" {
		data.Kewarganegaraan = sql.NullString{
			String: payload.Kewarganegaraan,
			Valid:  true,
		}
	}
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
	data.StatusMenikah = sql.NullBool{
		Bool:  payload.StatusMenikah,
		Valid: true,
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

	if payload.AsalKota != "" {
		data.AsalKota = sql.NullString{
			String: payload.AsalKota,
			Valid:  true,
		}
	}

	if payload.GajiTerakhir != 0 {
		data.GajiTerakhir = sql.NullString{
			String: fmt.Sprintf("%.f", payload.GajiTerakhir),
			Valid:  true,
		}
	}

	if payload.HarapanGaji != 0 {
		data.HarapanGaji = sql.NullString{
			String: fmt.Sprintf("%.f", payload.HarapanGaji),
			Valid:  true,
		}
	}
	if payload.SedangBekerja != "" {
		data.SedangBekerja = sql.NullString{
			String: payload.SedangBekerja,
			Valid:  true,
		}
	}

	if payload.KetersediaanBekerja != "" {
		parsedTime, err := time.Parse("2006-01-02", payload.KetersediaanBekerja)
		if err == nil {
			data.KetersediaanBekerja = sql.NullTime{Time: parsedTime, Valid: true}
		}
	}

	if payload.SumberInformasi != "" {
		data.SumberInformasi = sql.NullString{String: payload.SumberInformasi, Valid: true}
	}

	if payload.Alasan != "" {
		data.Alasan = sql.NullString{
			String: payload.Alasan,
			Valid:  true,
		}
	}

	if payload.KetersediaanInter != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", payload.KetersediaanInter)
		if err == nil {
			data.KetersediaanInter = sql.NullTime{Time: parsedTime, Valid: true}
		}
	}
	return
}
