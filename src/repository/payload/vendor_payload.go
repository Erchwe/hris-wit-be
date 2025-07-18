package payload

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type CreateVendorPayload struct {
	VendorID   string `json:"vendor_id"`
	NamaVendor string `json:"nama_vendor" valid:"required"`
	Alamat     string `json:"alamat" valid:"required"`
	Status     string `json:"status" valid:"required"`
}

func (p *CreateVendorPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return nil
}

func (p *CreateVendorPayload) ToEntity(user sqlc.GetUserBackofficeRow) sqlc.CreateVendorParams {
	var (
		id = utility.GenerateGoogleUUID()
	)

	return sqlc.CreateVendorParams{
		VendorID:   id,
		NamaVendor: p.NamaVendor,
		Alamat:     p.Alamat,
		Status:     p.Status,
		CreatedBy:  user.CreatedBy,
	}
}
