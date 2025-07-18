package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type UpdateKategoriPayload struct {
	KategoriID   string `json:"kategori_id" valid:"required"`
	NamaKategori string `json:"nama_kategori" valid:"required"`
	UpdatedBy    string `json:"updated_by" valid:"required"`
}

func (p *UpdateKategoriPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *UpdateKategoriPayload) ToEntity() sqlc.UpdateKategoriParams {
	return sqlc.UpdateKategoriParams{
		KategoriID:   p.KategoriID,
		NamaKategori: p.NamaKategori,
		UpdatedBy:    sql.NullString{String: p.UpdatedBy, Valid: true},
	}
}
