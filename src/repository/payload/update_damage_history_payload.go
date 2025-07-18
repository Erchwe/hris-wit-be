package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"time"
)

type UpdateDamageHistoryPayload struct {
	DamageHistoryID     string `json:"damage_history_id" valid:"required"`
	InventarisID        string `json:"inventaris_id" valid:"required"`
	IDPegawai           string `json:"id_pegawai" valid:"required"`
	TglRusak            string `json:"tgl_rusak" valid:"required"`
	TglAwalPerbaikan    string `json:"tgl_awal_perbaikan" valid:"required"`
	TglSelesaiPerbaikan string `json:"tgl_selesai_perbaikan" valid:"required"`
	Description         string `json:"description"`
	BiayaPerbaikan      int64  `json:"biaya_perbaikan" valid:"required"`
	VendorPerbaikan     int64  `json:"vendor_perbaikan" valid:"required"`
	Status              string `json:"status" valid:"required"`
	UpdatedBy           string `json:"updated_by" valid:"required"`
}

func (p *UpdateDamageHistoryPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *UpdateDamageHistoryPayload) ToEntity() (sqlc.UpdateDamageHistoryParams, error) {
	parse := func(dateStr string) (time.Time, error) {
		return time.Parse("2006-01-02", dateStr)
	}

	tglRusak, err := parse(p.TglRusak)
	if err != nil {
		return sqlc.UpdateDamageHistoryParams{}, errors.Wrap(httpservice.ErrBadRequest, "invalid tgl_rusak")
	}
	tglAwal, err := parse(p.TglAwalPerbaikan)
	if err != nil {
		return sqlc.UpdateDamageHistoryParams{}, errors.Wrap(httpservice.ErrBadRequest, "invalid tgl_awal_perbaikan")
	}
	tglSelesai, err := parse(p.TglSelesaiPerbaikan)
	if err != nil {
		return sqlc.UpdateDamageHistoryParams{}, errors.Wrap(httpservice.ErrBadRequest, "invalid tgl_selesai_perbaikan")
	}

	return sqlc.UpdateDamageHistoryParams{
		DamageHistoryID:     p.DamageHistoryID,
		InventarisID:        p.InventarisID,
		IDPegawai:           p.IDPegawai,
		TglRusak:            tglRusak,
		TglAwalPerbaikan:    tglAwal,
		TglSelesaiPerbaikan: tglSelesai,
		Description:         sql.NullString{String: p.Description, Valid: p.Description != ""},
		BiayaPerbaikan:      p.BiayaPerbaikan,
		VendorPerbaikan:     p.VendorPerbaikan,
		Status:              p.Status,
		UpdatedBy:           sql.NullString{String: p.UpdatedBy, Valid: true},
	}, nil
}
