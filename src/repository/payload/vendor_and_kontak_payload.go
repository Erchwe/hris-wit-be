package payload

import (
	"github.com/pkg/errors"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type CreateVendorWithKontakPayload struct {
	Vendor CreateVendorPayload         `json:"vendor"`
	Kontak []CreateKontakVendorPayload `json:"kontak"`
}

func (p *CreateVendorWithKontakPayload) Validate() error {
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

func (p *CreateVendorWithKontakPayload) ToEntities(cfg config.KVStore, user sqlc.GetUserBackofficeRow) (sqlc.CreateVendorParams, []sqlc.CreateKontakVendorParams) {
	vendorParams := p.Vendor.ToEntity(user)

	var kontakParams []sqlc.CreateKontakVendorParams
	for _, k := range p.Kontak {
		k.VendorID = vendorParams.VendorID
		kontakParams = append(kontakParams, k.ToEntity(user))
	}

	return vendorParams, kontakParams
}
