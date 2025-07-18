package payload

import (
	"database/sql"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type CreateRuanganPayload struct {
	RuanganID   string `json:"ruangan_id"`
	KantorID    string `json:"kantor_id" valid:"required"`
	NamaRuangan string `json:"nama_ruangan" valid:"required"`
	Lantai      string `json:"lantai"` // string supaya bisa dikosongkan
	Status      string `json:"status" valid:"required"`
}

func (p *CreateRuanganPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrapf(httpservice.ErrBadRequest, "bad request: %s", err.Error())
	}
	return nil
}

func (p *CreateRuanganPayload) ToEntity(cfg config.KVStore, user sqlc.GetUserBackofficeRow) (data sqlc.CreateRuanganParams, err error) {
	id := p.RuanganID
	if id == "" {
		id = utility.GenerateGoogleUUID()
	}

	data = sqlc.CreateRuanganParams{
		RuanganID:   id,
		KantorID:    p.KantorID,
		NamaRuangan: p.NamaRuangan,
		Status:      p.Status,
		CreatedBy:   user.CreatedBy,
	}

	if p.Lantai != "" {
		l, parseErr := strconv.Atoi(p.Lantai)
		if parseErr == nil {
			data.Lantai = sql.NullInt32{Int32: int32(l), Valid: true}
		}
	}

	return
}
