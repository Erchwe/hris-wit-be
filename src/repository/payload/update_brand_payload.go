package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type UpdateBrandPayload struct {
	BrandID   string `json:"brand_id" valid:"required"`
	NamaBrand string `json:"nama_brand" valid:"required"`
	Status    string `json:"status" valid:"required"`
	UpdatedBy string `json:"updated_by" valid:"required"`
}

func (p *UpdateBrandPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *UpdateBrandPayload) ToEntity() sqlc.UpdateBrandParams {
	return sqlc.UpdateBrandParams{
		BrandID:   p.BrandID,
		NamaBrand: p.NamaBrand,
		Status:    p.Status,
		UpdatedBy: sql.NullString{String: p.UpdatedBy, Valid: true},
	}
}
