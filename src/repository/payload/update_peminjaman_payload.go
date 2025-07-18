package payload

import (
	"database/sql"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type UpdatePeminjamanPayload struct {
	// PeminjamanID datang dari URL parameter, tidak perlu di payload body request
	InventarisID     string `json:"inventaris_id" valid:"required"` // Asumsi ini juga bisa diupdate
	TglPinjam        string `json:"tgl_pinjam" valid:"required"`
	TglKembali       string `json:"tgl_kembali" valid:"required"`
	StatusPeminjaman string `json:"status_peminjaman" valid:"required"`
	Notes            string `json:"notes"`
	// UpdatedBy tidak perlu di payload request, akan diambil dari JWT context
}

func (p *UpdatePeminjamanPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *UpdatePeminjamanPayload) ToEntity(user sqlc.GetUserBackofficeRow) (data sqlc.UpdatePeminjamanParams, err error) { // <<-- Tambah parameter 'user'
	tglPinjam, err := time.Parse("2006-01-02", p.TglPinjam)
	if err != nil {
		err = errors.Wrap(httpservice.ErrBadRequest, "invalid format for tgl_pinjam")
		return data, err
	}
	tglKembali, err := time.Parse("2006-01-02", p.TglKembali)
	if err != nil {
		err = errors.Wrap(httpservice.ErrBadRequest, "invalid format for tgl_kembali")
		return data, err
	}

	updatedByName := ""
	if user.Name.Valid {
		updatedByName = user.Name.String
	} else {
		updatedByName = "system_user"
	}

	data = sqlc.UpdatePeminjamanParams{
		TglPinjam:        tglPinjam,
		TglKembali:       tglKembali,
		StatusPeminjaman: sqlc.StatusPeminjamanEnum(p.StatusPeminjaman),
		Notes:            sql.NullString{String: p.Notes, Valid: p.Notes != ""},
		UpdatedBy:        sql.NullString{String: updatedByName, Valid: true}, // Gunakan updatedByName dari user context
	}
	return data, nil // Pastikan return data dan nil
}
