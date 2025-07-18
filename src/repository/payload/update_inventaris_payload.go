package payload

import (
	"database/sql"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type UpdateInventarisPayload struct {
	InventarisID     string `json:"inventaris_id" form:"inventaris_id" valid:"required"`
	BrandID          string `json:"brand_id" form:"brand_id" valid:"required"`
	RuanganID        string `json:"ruangan_id" form:"ruangan_id"`
	UserID           string `json:"user_id" form:"user_id"`
	KategoriID       string `json:"kategori_id" form:"kategori_id" valid:"required"`
	VendorID         string `json:"vendor_id" form:"vendor_id" valid:"required"`
	NamaInventaris   string `json:"nama_inventaris" form:"nama_inventaris" valid:"required"`
	Jumlah           int    `json:"jumlah" form:"jumlah" valid:"required"`
	TanggalBeli      string `json:"tanggal_beli" form:"tanggal_beli" valid:"required"`
	Harga            int64  `json:"harga" form:"harga" valid:"required"`
	Keterangan       string `json:"keterangan" form:"keterangan"`
	OldInventoryCode string `json:"old_inventory_code" form:"old_inventory_code"`
	ImageURL         string `json:"image_url" form:"image_url"`
	Status           string `json:"status" form:"status" valid:"required"`
	UpdatedBy        string `json:"updated_by" form:"updated_by" valid:"required"`
}

func (p *UpdateInventarisPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(err, "invalid update payload")
	}
	return nil
}

func (p *UpdateInventarisPayload) ToEntity() sqlc.UpdateInventarisParams {
	tglBeli, _ := time.Parse("2006-01-02", p.TanggalBeli)

	return sqlc.UpdateInventarisParams{
		BrandID:          p.BrandID,
		RuanganID:        sql.NullString{String: p.RuanganID, Valid: p.RuanganID != ""},
		UserID:           sql.NullString{String: p.UserID, Valid: p.UserID != ""},
		KategoriID:       p.KategoriID,
		VendorID:         p.VendorID,
		NamaInventaris:   p.NamaInventaris,
		Jumlah:           int32(p.Jumlah),
		TanggalBeli:      tglBeli,
		Harga:            p.Harga,
		Keterangan:       sql.NullString{String: p.Keterangan, Valid: p.Keterangan != ""},
		OldInventoryCode: sql.NullString{String: p.OldInventoryCode, Valid: p.OldInventoryCode != ""},
		ImageUrl:         sql.NullString{String: p.ImageURL, Valid: p.ImageURL != ""},
		Status:           p.Status,
		UpdatedBy:        sql.NullString{String: p.UpdatedBy, Valid: true},
		InventarisID:     p.InventarisID,
	}
}
