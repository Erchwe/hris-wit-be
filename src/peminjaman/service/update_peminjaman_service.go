package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/common/helper"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func parseStatusPeminjamanEnumManual(s string) (sqlc.StatusPeminjamanEnum, error) {
	switch s {
	case "Menunggu Persetujuan":
		return sqlc.StatusPeminjamanEnumMenungguPersetujuan, nil
	case "Sedang Dipinjam":
		return sqlc.StatusPeminjamanEnumSedangDipinjam, nil
	case "Tidak Dipinjam":
		return sqlc.StatusPeminjamanEnumTidakDipinjam, nil
	default:
		return "", fmt.Errorf("invalid StatusPeminjamanEnum: %q", s)
	}
}

func (s *PeminjamanService) UpdatePeminjaman(
	ctx context.Context,
	peminjamanID string,
	request payload.UpdatePeminjamanPayload, // <<--- GANTI TIPE REQUEST
	user sqlc.GetUserBackofficeRow,
) (payload.PeminjamanResponse, error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx for update peminjaman")
		return payload.PeminjamanResponse{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback update peminjaman", rollBackErr)
			}
		}
	}()

	// Panggil ToEntity dari request payload untuk mendapatkan sqlc.UpdatePeminjamanParams
	entity, err := request.ToEntity(user) // <<--- PANGGIL ToEntity dan teruskan 'user'
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to convert payload to peminjaman entity for update")
		return payload.PeminjamanResponse{}, errors.WithStack(httpservice.ErrBadRequest)
	}

	// Pastikan PeminjamanID di entity sesuai dengan ID dari URL
	// Karena PeminjamanID di payload tidak digunakan, kita harus set dari path param
	entity.PeminjamanID = peminjamanID

	updatedPeminjaman, err := q.UpdatePeminjaman(ctx, entity) // <<--- TERUSKAN ENTITY LANGSUNG
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed update peminjaman: %v", err)
		return payload.PeminjamanResponse{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	result := payload.PeminjamanResponse{
		PeminjamanID:     updatedPeminjaman.PeminjamanID,
		TanggalPinjam:    updatedPeminjaman.TglPinjam,
		TanggalKembali:   updatedPeminjaman.TglKembali,
		StatusPeminjaman: string(updatedPeminjaman.StatusPeminjaman),
		Notes:            updatedPeminjaman.Notes.String,
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit update peminjaman")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}

func (s *PeminjamanService) UpdatePeminjamanStatus(
	ctx context.Context,
	peminjamanID string,
	request payload.UpdatePeminjamanStatusPayload,
	user sqlc.GetUserBackofficeRow,
) (payload.PeminjamanResponse, error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx for update peminjaman status")
		return payload.PeminjamanResponse{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback update peminjaman status", rollBackErr)
			}
		}
	}()

	statusEnum, err := helper.ParseStatusPeminjamanEnumManual(request.StatusPeminjaman)
	if err != nil {
		log.FromCtx(ctx).Error(err, "invalid status peminjaman enum value", request.StatusPeminjaman)
		return payload.PeminjamanResponse{}, errors.WithStack(httpservice.ErrBadRequest)
	}

	updatedPeminjaman, err := q.UpdatePeminjamanStatus(ctx, sqlc.UpdatePeminjamanStatusParams{
		PeminjamanID:     peminjamanID,
		StatusPeminjaman: statusEnum,
		UpdatedBy:        sql.NullString{String: user.Name.String, Valid: user.Name.Valid},
	})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed update peminjaman status: %v", err)
		return payload.PeminjamanResponse{}, errors.WithStack(httpservice.ErrUnknownSource)
	}

	result := payload.PeminjamanResponse{
		PeminjamanID:     updatedPeminjaman.PeminjamanID,
		TanggalPinjam:    updatedPeminjaman.TglPinjam,
		TanggalKembali:   updatedPeminjaman.TglKembali,
		StatusPeminjaman: string(updatedPeminjaman.StatusPeminjaman),
		Notes:            updatedPeminjaman.Notes.String,
	}

	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit update peminjaman status")
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	return result, nil
}
