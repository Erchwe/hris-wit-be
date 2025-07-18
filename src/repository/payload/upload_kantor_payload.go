package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type UpdateKantorPayload struct {
	KantorID   string `json:"kantor_id" valid:"required"`
	NamaKantor string `json:"nama_kantor" valid:"required"`
	Kota       string `json:"kota" valid:"required"`
	Alamat     string `json:"alamat" valid:"required"`
	NomorTelp  string `json:"nomor_telp" valid:"required"`
	Status     string `json:"status" valid:"required"`
	UpdatedBy  string `json:"updated_by" valid:"required"`
}

func (p *UpdateKantorPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *UpdateKantorPayload) ToEntity() sqlc.UpdateKantorParams {
	return sqlc.UpdateKantorParams{
		KantorID:   p.KantorID,
		NamaKantor: p.NamaKantor,
		Kota:       p.Kota,
		Alamat:     p.Alamat,
		NomorTelp:  p.NomorTelp,
		Status:     p.Status,
		UpdatedBy:  sql.NullString{String: p.UpdatedBy, Valid: true},
	}
}
