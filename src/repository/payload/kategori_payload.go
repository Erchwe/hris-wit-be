package payload

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type CreateKategoriPayload struct {
	KategoriID   string `json:"kategori_id"`
	NamaKategori string `json:"nama_kategori" valid:"required"`
}

func (p *CreateKategoriPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *CreateKategoriPayload) ToEntity(cfg config.KVStore, user sqlc.GetUserBackofficeRow) sqlc.CreateKategoriParams {
	id := p.KategoriID
	if id == "" {
		id = utility.GenerateGoogleUUID()
	}

	return sqlc.CreateKategoriParams{
		KategoriID:   id,
		NamaKategori: p.NamaKategori,
		CreatedBy:    user.CreatedBy,
	}
}
