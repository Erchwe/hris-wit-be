package payload

import (
	"database/sql"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type CreateKontakVendorPayload struct {
	KontakVendorID string `json:"kontak_vendor_id"`
	VendorID       string `json:"vendor_id"`
	JenisKontak    string `json:"jenis_kontak" valid:"required"`
	IsiKontak      string `json:"isi_kontak" valid:"required"`
	IsPrimary      *bool  `json:"is_primary"`
}

func (p *CreateKontakVendorPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return nil
}

func (p *CreateKontakVendorPayload) ToEntity(user sqlc.GetUserBackofficeRow) sqlc.CreateKontakVendorParams {
	id := p.KontakVendorID
	if id == "" {
		id = utility.GenerateGoogleUUID()
	}

	isPrimary := false
	if p.IsPrimary != nil {
		isPrimary = *p.IsPrimary
	}

	return sqlc.CreateKontakVendorParams{
		KontakVendorID: id,
		VendorID:       p.VendorID,
		JenisKontak:    p.JenisKontak,
		IsiKontak:      sql.NullString{String: p.IsiKontak, Valid: p.IsiKontak != ""},
		IsPrimary:      sql.NullBool{Bool: isPrimary, Valid: true},
		CreatedBy:      user.CreatedBy,
	}
}
