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

type InsertPelamarPengalamanKerjaPayload struct {
	// IdPengalamanKerja string  `json:"id_pengalaman_kerja"`
	IdPelamar      string `json:"id_pelamar"`
	NamaPerusahaan string `json:"nama_perusahaan"`
	Periode        string `json:"periode"`
	Jabatan        string `json:"jabatan"`
	Gaji           string `json:"gaji"`
	AlasanPindah   string `json:"alasan_pindah"`
}

func (payload *InsertPelamarPengalamanKerjaPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return nil
}

func (payload *InsertPelamarPengalamanKerjaPayload) ToEntity(cfg config.KVStore) (data sqlc.CreatePelamarPengalamanKerjaParams) {
	var (
		userId = utility.GenerateGoogleUUID()
	)

	data = sqlc.CreatePelamarPengalamanKerjaParams{
		IDPengalamanKerja: userId,
		IDPelamar:         payload.IdPelamar,
		NamaPerusahaan:    payload.NamaPerusahaan,
		Periode:           sql.NullString{},
		Jabatan:           sql.NullString{},
		Gaji:              sql.NullString{},
		AlasanPindah:      sql.NullString{},
		CreatedBy:         userId,
	}
	if payload.Periode != "" {
		data.Periode = sql.NullString{
			String: payload.Periode,
			Valid:  true,
		}
	}

	if payload.Jabatan != "" {
		data.Jabatan = sql.NullString{
			String: payload.Jabatan,
			Valid:  true,
		}
	}

	if payload.Gaji != "" {
		data.Gaji = sql.NullString{
			String: payload.Gaji,
			Valid:  true,
		}
	}

	if payload.AlasanPindah != "" {
		data.AlasanPindah = sql.NullString{
			String: payload.AlasanPindah,
			Valid:  true,
		}
	}
	return
}
