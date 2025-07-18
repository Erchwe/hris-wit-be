package payload

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type CreateBrandPayload struct {
	BrandID   string `json:"brand_id"` // opsional (generate jika kosong)
	NamaBrand string `json:"nama_brand" valid:"required"`
	Status    string `json:"status" valid:"required"`
}

func (p *CreateBrandPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *CreateBrandPayload) ToEntity(user sqlc.GetUserBackofficeRow) sqlc.CreateBrandParams {
	id := p.BrandID
	if id == "" {
		id = utility.GenerateGoogleUUID()
	}

	return sqlc.CreateBrandParams{
		BrandID:   id,
		NamaBrand: p.NamaBrand,
		Status:    p.Status,
		CreatedBy: user.CreatedBy,
	}
}
