package payload

import (
	"database/sql"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type CreateInventarisPayload struct {
	InventarisID     string `form:"inventaris_id"` // optional
	BrandID          string `form:"brand_id" valid:"required"`
	RuanganID        string `form:"ruangan_id"`
	UserID           string `form:"user_id"`
	KategoriID       string `form:"kategori_id" valid:"required"`
	VendorID         string `form:"vendor_id" valid:"required"`
	NamaInventaris   string `form:"nama_inventaris" valid:"required"`
	Jumlah           int    `form:"jumlah" valid:"required"`
	TanggalBeli      string `form:"tanggal_beli" valid:"required"` // YYYY-MM-DD
	Harga            int64  `form:"harga" valid:"required"`
	Keterangan       string `form:"keterangan"`
	OldInventoryCode string `form:"old_inventory_code"`
	ImageURL         string `form:"image_url"`
	Status           string `form:"status" valid:"required"`
}

func (p *CreateInventarisPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return nil
}

func (p *CreateInventarisPayload) ToEntity(user sqlc.GetUserBackofficeRow) (sqlc.CreateInventarisParams, error) {
	tglBeli, err := time.Parse("2006-01-02", p.TanggalBeli)
	if err != nil {
		return sqlc.CreateInventarisParams{}, errors.Wrap(httpservice.ErrBadRequest, "invalid tanggal_beli")
	}

	id := p.InventarisID
	if id == "" {
		id = utility.GenerateGoogleUUID()
	}

	return sqlc.CreateInventarisParams{
		InventarisID:     id,
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
		CreatedBy:        user.CreatedBy,
	}, nil
}
