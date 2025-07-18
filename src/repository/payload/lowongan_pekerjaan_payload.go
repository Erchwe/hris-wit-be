package payload

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type InsertLowonganPekerjaanPayload struct {
	//IdLowonganPekerjaan string `json:"id_lowongan_pekerjaan" valid:"required"`
	Posisi           string `json:"posisi" valid:"required"`
	TglBukaLowongan  string `json:"tgl_buka_lowongan" valid:"required"`
	TglTutupLowongan string `json:"tgl_tutup_lowongan" valid:"required"`
	Kriteria         string `json:"kriteria"`
	Deskripsi        string `json:"deskripsi"`
	LinkLowongan     string `json:"link_lowongan"`
}

type ListLowonganPekerjaanPayload struct {
	IDLowonganPekerjaan string         `json:"id_lowongan_pekerjaan"`
	Posisi              string         `json:"posisi"`
	TglBukaLowongan     string         `json:"tgl_buka_lowongan"`
	TglTutupLowongan    string         `json:"tgl_tutup_lowongan"`
	LinkLowongan        sql.NullString `json:"link_lowongan"`
}

type UpdateLowonganPekerjaanPayload struct {
	IDLowonganPekerjaan string `json:"id_lowongan_pekerjaan" valid:"required"`
	Posisi              string `json:"posisi"`
	TglBukaLowongan     string `json:"tgl_buka_lowongan"`
	TglTutupLowongan    string `json:"tgl_tutup_lowongan"`
	Kriteria            string `json:"kriteria"`
	Deskripsi           string `json:"deskripsi"`
	LinkLowongan        string `json:"link_lowongan"`
}

func (payload *UpdateLowonganPekerjaanPayload) Validate() (err error) {
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
		return
	}
	return
}

func (payload *InsertLowonganPekerjaanPayload) Validate() (err error) {
	// Validasi required fields
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		return errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}

	// Validasi tanggal buka <= tanggal tutup
	tglBuka, err := time.Parse("2006-01-02", payload.TglBukaLowongan)
	if err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, "format tanggal buka tidak valid (YYYY-MM-DD)")
	}

	tglTutup, err := time.Parse("2006-01-02", payload.TglTutupLowongan)
	if err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, "format tanggal tutup tidak valid (YYYY-MM-DD)")
	}

	if tglBuka.After(tglTutup) {
		return errors.Wrap(httpservice.ErrBadRequest, "tanggal buka tidak boleh lebih besar dari tanggal tutup")
	}

	return nil
}

func (payload *InsertLowonganPekerjaanPayload) ToEntity(cfg config.KVStore) (data sqlc.CreateLowonganPekerjaanParams, err error) {
	var (
		userId = utility.GenerateGoogleUUID()
	)

	tglBuka, err := time.Parse("2006-01-02", payload.TglBukaLowongan)
	if err != nil {
		return data, errors.Wrapf(httpservice.ErrBadRequest, "invalid tgl_buka_lowongan format")
	}
	tglTutup, err := time.Parse("2006-01-02", payload.TglTutupLowongan)
	if err != nil {
		return data, errors.Wrapf(httpservice.ErrBadRequest, "invalid tgl_tutup_lowongan format")
	}

	linkLowongan := fmt.Sprintf("http://localhost:5173/pelamarForm?id=%s", userId)

	data = sqlc.CreateLowonganPekerjaanParams{
		IDLowonganPekerjaan: userId,
		Posisi:              payload.Posisi,
		TglBukaLowongan:     tglBuka,
		TglTutupLowongan:    tglTutup,
		Kriteria:            sql.NullString{},
		Deskripsi:           sql.NullString{},
		LinkLowongan:        sql.NullString{String: linkLowongan, Valid: true},
		CreatedBy:           userId, //guid dari token siapa yang login
	}

	if payload.Kriteria != "" {
		data.Kriteria = sql.NullString{
			String: payload.Kriteria,
			Valid:  true,
		}
	}
	if payload.Deskripsi != "" {
		data.Deskripsi = sql.NullString{
			String: payload.Deskripsi,
			Valid:  true,
		}
	}

	return
}

func ListLowonganPekerjaan(payload sqlc.LowonganPekerjaan) ListLowonganPekerjaanPayload {
	tglBuka := payload.TglBukaLowongan.Format("2006-01-02")
	tglTutup := payload.TglTutupLowongan.Format("2006-01-02")

	return ListLowonganPekerjaanPayload{
		IDLowonganPekerjaan: payload.IDLowonganPekerjaan,
		Posisi:              payload.Posisi,
		TglBukaLowongan:     tglBuka,
		TglTutupLowongan:    tglTutup,
		LinkLowongan:        payload.LinkLowongan,
	}
}

func (payload *UpdateLowonganPekerjaanPayload) ToEntity(old sqlc.LowonganPekerjaan) (data sqlc.UpdateLowonganPekerjaanParams, err error) {
	var tglBuka, tglTutup time.Time

	if payload.TglBukaLowongan != "" {
		tglBuka, err = time.Parse("2006-01-02", payload.TglBukaLowongan)
		if err != nil {
			return data, errors.Wrapf(httpservice.ErrBadRequest, "invalid tgl_buka_lowongan format")
		}
	} else {
		tglBuka = old.TglBukaLowongan
	}

	if payload.TglTutupLowongan != "" {
		tglTutup, err = time.Parse("2006-01-02", payload.TglTutupLowongan)
		if err != nil {
			return data, errors.Wrapf(httpservice.ErrBadRequest, "invalid tgl_tutup_lowongan format")
		}
	} else {
		tglTutup = old.TglTutupLowongan
	}

	data = sqlc.UpdateLowonganPekerjaanParams{
		IDLowonganPekerjaan: payload.IDLowonganPekerjaan,
		Posisi:              payload.Posisi,
		TglBukaLowongan:     tglBuka,
		TglTutupLowongan:    tglTutup,
		Kriteria:            sql.NullString{Valid: false},
		Deskripsi:           sql.NullString{Valid: false},
		LinkLowongan:        sql.NullString{Valid: false},
	}

	if payload.Kriteria != "" {
		data.Kriteria = sql.NullString{String: payload.Kriteria, Valid: true}
	}
	if payload.Deskripsi != "" {
		data.Deskripsi = sql.NullString{String: payload.Deskripsi, Valid: true}
	}
	if payload.LinkLowongan != "" {
		data.LinkLowongan = sql.NullString{String: payload.LinkLowongan, Valid: true}
	}

	return
}
