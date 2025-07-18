package payload

import (
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/common/utility"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type UpdateVendorPayload struct {
	VendorID   string `json:"vendor_id" valid:"required"`
	NamaVendor string `json:"nama_vendor"`
	Alamat     string `json:"alamat"`
	Status     string `json:"status"`
	UpdatedBy  string `json:"updated_by" valid:"required"`
}

type UpdateVendorWithKontakPayload struct {
	Vendor UpdateVendorPayload         `json:"vendor"`
	Kontak []UpdateKontakVendorPayload `json:"kontak"`
}

func (p *UpdateVendorPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(err, "invalid update vendor payload")
	}
	return nil
}

func (p *UpdateVendorPayload) ToEntity() sqlc.UpdateVendorParams {
	return sqlc.UpdateVendorParams{
		NamaVendor: p.NamaVendor,
		Alamat:     p.Alamat,
		Status:     p.Status,
		UpdatedBy:  sql.NullString{String: p.UpdatedBy, Valid: true},
		VendorID:   p.VendorID,
	}
}

func (p *UpdateVendorWithKontakPayload) Validate() error {
	if err := p.Vendor.Validate(); err != nil {
		return err
	}
	for i, k := range p.Kontak {
		if err := k.Validate(); err != nil {
			return errors.Wrapf(err, "kontak index %d invalid", i)
		}
	}
	return nil
}

func (p *UpdateVendorWithKontakPayload) ToEntities(user sqlc.GetUserBackofficeRow) (sqlc.UpdateVendorParams, []sqlc.CreateKontakVendorParams) {
	vendor := p.Vendor.ToEntity()

	var kontakList []sqlc.CreateKontakVendorParams
	for _, k := range p.Kontak {
		isPrimary := false
		if k.IsPrimary != nil {
			isPrimary = *k.IsPrimary
		}

		kontakList = append(kontakList, sqlc.CreateKontakVendorParams{
			KontakVendorID: utility.GenerateGoogleUUID(),
			VendorID:       p.Vendor.VendorID,
			JenisKontak:    k.JenisKontak,
			IsiKontak:      sql.NullString{String: k.IsiKontak, Valid: k.IsiKontak != ""},
			IsPrimary:      sql.NullBool{Bool: isPrimary, Valid: true},
			CreatedBy:      user.UpdatedBy.String,
		})
	}

	return vendor, kontakList
}
