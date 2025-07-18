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

type CreateDamageHistoryPayload struct {
	DamageHistoryID     string `json:"damage_history_id"`
	InventarisID        string `json:"inventaris_id" valid:"required"`
	IDPegawai           string `json:"id_pegawai" valid:"required"`
	TglRusak            string `json:"tgl_rusak" valid:"required"`
	TglAwalPerbaikan    string `json:"tgl_awal_perbaikan" valid:"required"`
	TglSelesaiPerbaikan string `json:"tgl_selesai_perbaikan" valid:"required"`
	Description         string `json:"description"`
	BiayaPerbaikan      int64  `json:"biaya_perbaikan" valid:"required"`
	VendorPerbaikan     int64  `json:"vendor_perbaikan" valid:"required"`
	Status              string `json:"status" valid:"required"`
}

func (p *CreateDamageHistoryPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *CreateDamageHistoryPayload) ToEntity(cfg config.KVStore, user sqlc.GetUserBackofficeRow) (sqlc.CreateDamageHistoryParams, error) {
	id := p.DamageHistoryID
	if id == "" {
		id = utility.GenerateGoogleUUID()
	}

	parse := func(dateStr string) (time.Time, error) {
		return time.Parse("2006-01-02", dateStr)
	}

	tglRusak, err := parse(p.TglRusak)
	if err != nil {
		return sqlc.CreateDamageHistoryParams{}, errors.Wrap(httpservice.ErrBadRequest, "invalid tgl_rusak")
	}
	tglAwal, err := parse(p.TglAwalPerbaikan)
	if err != nil {
		return sqlc.CreateDamageHistoryParams{}, errors.Wrap(httpservice.ErrBadRequest, "invalid tgl_awal_perbaikan")
	}
	tglSelesai, err := parse(p.TglSelesaiPerbaikan)
	if err != nil {
		return sqlc.CreateDamageHistoryParams{}, errors.Wrap(httpservice.ErrBadRequest, "invalid tgl_selesai_perbaikan")
	}

	return sqlc.CreateDamageHistoryParams{
		DamageHistoryID:     id,
		InventarisID:        p.InventarisID,
		IDPegawai:           p.IDPegawai,
		TglRusak:            tglRusak,
		TglAwalPerbaikan:    tglAwal,
		TglSelesaiPerbaikan: tglSelesai,
		Description:         sql.NullString{String: p.Description, Valid: p.Description != ""},
		BiayaPerbaikan:      p.BiayaPerbaikan,
		VendorPerbaikan:     p.VendorPerbaikan,
		Status:              p.Status,
		CreatedBy:           user.CreatedBy,
	}, nil
}
