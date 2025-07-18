package payload

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type CreateKantorPayload struct {
	KantorID   string `json:"kantor_id"` // opsional
	NamaKantor string `json:"nama_kantor" valid:"required"`
	Kota       string `json:"kota" valid:"required"`
	Alamat     string `json:"alamat" valid:"required"`
	NomorTelp  string `json:"nomor_telp" valid:"required"`
	Status     string `json:"status" valid:"required"`
}

func (p *CreateKantorPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *CreateKantorPayload) ToEntity(cfg config.KVStore, user sqlc.GetUserBackofficeRow) sqlc.CreateKantorParams {
	id := p.KantorID
	if id == "" {
		id = utility.GenerateGoogleUUID()
	}

	return sqlc.CreateKantorParams{
		KantorID:   id,
		NamaKantor: p.NamaKantor,
		Kota:       p.Kota,
		Alamat:     p.Alamat,
		NomorTelp:  p.NomorTelp,
		Status:     p.Status,
		CreatedBy:  user.CreatedBy,
	}
}
