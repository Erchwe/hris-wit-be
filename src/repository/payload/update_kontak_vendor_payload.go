package payload

import (
	"database/sql"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type UpdateKontakVendorPayload struct {
	KontakVendorID string `json:"kontak_vendor_id" valid:"required"`
	VendorID       string `json:"vendor_id"`
	JenisKontak    string `json:"jenis_kontak" valid:"required"`
	IsiKontak      string `json:"isi_kontak" valid:"required"`
	IsPrimary      *bool  `json:"is_primary"`
	UpdatedBy      string `json:"updated_by" valid:"required"`
}

func (p *UpdateKontakVendorPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return nil
}

func (p *UpdateKontakVendorPayload) ToEntity() sqlc.UpdateKontakVendorParams {
	isPrimary := false
	if p.IsPrimary != nil {
		isPrimary = *p.IsPrimary
	}

	return sqlc.UpdateKontakVendorParams{
		KontakVendorID: p.KontakVendorID,
		JenisKontak:    p.JenisKontak,
		IsiKontak:      sql.NullString{String: p.IsiKontak, Valid: p.IsiKontak != ""},
		IsPrimary:      sql.NullBool{Bool: isPrimary, Valid: true},
		UpdatedBy:      sql.NullString{String: p.UpdatedBy, Valid: true},
	}
}
