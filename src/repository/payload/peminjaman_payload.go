package payload

import (
	"database/sql"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

type UserAuthContext struct {
	GUID  string
	Name  string
	Email string
}

type CreatePeminjamanPayload struct {
	InventarisID string `json:"inventaris_id" valid:"required"`
	TglPinjam    string `json:"tgl_pinjam" valid:"required"`
	TglKembali   string `json:"tgl_kembali" valid:"required"`
	Notes        string `json:"notes"`
}

func (p *CreatePeminjamanPayload) Validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return errors.Wrap(httpservice.ErrBadRequest, err.Error())
	}
	return nil
}

func (p *CreatePeminjamanPayload) ToEntity(user sqlc.GetUserBackofficeRow) (data sqlc.CreatePeminjamanParams, err error) {
	id := utility.GenerateGoogleUUID()

	tglPinjam, err := time.Parse("2006-01-02", p.TglPinjam)
	if err != nil {
		return data, errors.Wrap(httpservice.ErrBadRequest, "invalid format for tgl_pinjam")
	}
	tglKembali, err := time.Parse("2006-01-02", p.TglKembali)
	if err != nil {
		return data, errors.Wrap(httpservice.ErrBadRequest, "invalid format for tgl_kembali")
	}

	createdByName := ""
	if user.Name.Valid {
		createdByName = user.Name.String
	}

	data = sqlc.CreatePeminjamanParams{
		PeminjamanID:     id,
		InventarisID:     p.InventarisID,
		UserID:           user.Guid,
		TglPinjam:        tglPinjam,
		TglKembali:       tglKembali,
		StatusPeminjaman: sqlc.StatusPeminjamanEnumMenungguPersetujuan,
		Notes:            sql.NullString{String: p.Notes, Valid: p.Notes != ""},
		CreatedBy:        createdByName,
	}
	return data, nil
}

// Payloads Respons
type PeminjamanResponse struct {
	PeminjamanID     string    `json:"peminjaman_id"`
	NamaInventaris   string    `json:"nama_inventaris"` // UBAH: string (akan di-handle di service)
	TanggalPinjam    time.Time `json:"tanggal_pinjam"`
	TanggalKembali   time.Time `json:"tanggal_kembali"`
	StatusPeminjaman string    `json:"status_peminjaman"`
	Notes            string    `json:"notes"` // UBAH: string (akan di-handle di service)
}

type PendingPeminjamanResponse struct {
	PeminjamanID   string         `json:"peminjaman_id"`
	NamaInventaris string         `json:"nama_inventaris"`
	NamaPeminjam   sql.NullString `json:"nama_peminjam"`
	UserIDPeminjam string         `json:"user_id_peminjam"`
	TanggalPinjam  time.Time      `json:"tanggal_pinjam"`
	TanggalKembali time.Time      `json:"tanggal_kembali"`
	Notes          sql.NullString `json:"notes"`
}

type UpdatePeminjamanStatusPayload struct {
	StatusPeminjaman string `json:"status_peminjaman" valid:"required"`
}

type AvailableInventarisResponse struct {
	InventarisID   string         `json:"inventaris_id"`
	NamaInventaris string         `json:"nama_inventaris"`
	Status         interface{}    `json:"status"` // TETAP interface{} karena itu yang dihasilkan sqlc
	Keterangan     sql.NullString `json:"keterangan"`
}

type OverduePeminjamanResponse struct {
	PeminjamanID     string    `json:"peminjaman_id"`
	NamaInventaris   string    `json:"nama_inventaris"`
	TanggalKembali   time.Time `json:"tanggal_kembali"`
	TanggalPinjam    time.Time `json:"tanggal_pinjam"`
	StatusPeminjaman string    `json:"status_peminjaman"`
}
