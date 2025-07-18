package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type UpdateDataPelamarPayload struct {
	IdDataPelamar   string  `json:"id_data_pelamar" valid:"required"`
	Email           string  `json:"email" valid:"email,required"`
	NamaLengkap     string  `json:"nama_lengkap" valid:"required"`
	TempatLahir     string  `json:"tempat_lahir"`
	TglLahir        string  `json:"tgl_lahir"`
	JenisKelamin    string  `json:"jenis_kelamin" valid:"required"`
	Kewarganegaraan string  `json:"kewarganegaraan"`
	Phone           string  `json:"phone"`
	Mobile          string  `json:"mobile"`
	Agama           string  `json:"agama"`
	GolDarah        string  `json:"gol_darah"`
	StatusMenikah   bool    `json:"status_menikah"`
	NoKTP           string  `json:"no_ktp"`
	NoNPWP          string  `json:"no_npwp"`
	Status          string  `json:"status" valid:"required,in(New|Short list|HR Interview|User Interview|Refference Checking|Offering|Psikotest|Hired|Rejected)"`
	AsalKota        string  `json:"asal_kota"`
	GajiTerakhir    float64 `json:"gaji_terakhir"`
	HarapanGaji     float64 `json:"harapan_gaji"`
	SedangBekerja   string  `json:"sedang_bekerja"`
	SumberInformasi string  `json:"sumber_informasi"`
	Alasan          string  `json:"alasan"`
	ProfesiKerja    string  `json:"profesi_kerja"`
}

func (p *UpdateDataPelamarPayload) Validate() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return errors.Wrap(err, "invalid payload data")
	}
	return nil
}

func (p *UpdateDataPelamarPayload) ToEntity(old sqlc.DataPelamar) sqlc.UpdatePelamarParams {
	data := sqlc.UpdatePelamarParams{
		Status: toNullString(p.Status),
	}

	return data
}

func toNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}
