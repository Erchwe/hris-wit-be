package payload

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"strconv"
)

type UpdateRuanganPayload struct {
	RuanganID   string `json:"ruangan_id" valid:"required"`
	KantorID    string `json:"kantor_id" valid:"required"`
	NamaRuangan string `json:"nama_ruangan" valid:"required"`
	Lantai      string `json:"lantai"` // tetap string
	Status      string `json:"status" valid:"required"`
	UpdatedBy   string `json:"updated_by" valid:"required"`
}

func (p *UpdateRuanganPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(err, "invalid update ruangan payload")
	}
	return nil
}

func (p *UpdateRuanganPayload) ToEntity() (data sqlc.UpdateRuanganParams, err error) {
	data = sqlc.UpdateRuanganParams{
		RuanganID:   p.RuanganID,
		KantorID:    p.KantorID,
		NamaRuangan: p.NamaRuangan,
		Status:      p.Status,
		UpdatedBy:   sql.NullString{String: p.UpdatedBy, Valid: true},
	}

	if p.Lantai != "" {
		l, parseErr := strconv.Atoi(p.Lantai)
		if parseErr == nil {
			data.Lantai = sql.NullInt32{Int32: int32(l), Valid: true}
		}
	}

	return
}
