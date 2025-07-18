package service

import (
	"context"
	"database/sql"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *PeminjamanService) GetPeminjamanByID(ctx context.Context, peminjamanID string) (payload.PeminjamanResponse, error) {
	q := sqlc.New(s.mainDB)
	var result payload.PeminjamanResponse

	peminjaman, err := q.GetPeminjamanByID(ctx, peminjamanID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get peminjaman by ID: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			return result, errors.WithStack(httpservice.ErrNoResultData)
		}
		return result, errors.WithStack(httpservice.ErrUnknownSource)
	}

	namaInventaris := ""
	if peminjaman.NamaInventaris.Valid {
		namaInventaris = peminjaman.NamaInventaris.String
	}

	result = payload.PeminjamanResponse{
		PeminjamanID:     peminjaman.PeminjamanID,
		NamaInventaris:   namaInventaris,
		TanggalPinjam:    peminjaman.TglPinjam,
		TanggalKembali:   peminjaman.TglKembali,
		StatusPeminjaman: string(peminjaman.StatusPeminjaman),
		Notes:            peminjaman.Notes.String,
	}

	return result, nil
}

func (s *PeminjamanService) GetListPeminjaman(ctx context.Context) ([]payload.PeminjamanResponse, error) {
	q := sqlc.New(s.mainDB)

	peminjamanList, err := q.ListPeminjaman(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list peminjaman from DB: %v", err)
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	results := make([]payload.PeminjamanResponse, len(peminjamanList))
	for i, p := range peminjamanList {

		namaInventaris := ""
		if p.NamaInventaris.Valid {
			namaInventaris = p.NamaInventaris.String
		}

		results[i] = payload.PeminjamanResponse{
			PeminjamanID:     p.PeminjamanID,
			NamaInventaris:   namaInventaris,
			TanggalPinjam:    p.TglPinjam,
			TanggalKembali:   p.TglKembali,
			StatusPeminjaman: string(p.StatusPeminjaman),
			Notes:            p.Notes.String,
		}
	}
	return results, nil
}

func (s *PeminjamanService) GetListPeminjamanByUserID(ctx context.Context, userID string) ([]payload.PeminjamanResponse, error) {
	q := sqlc.New(s.mainDB)
	peminjamanList, err := q.ListPeminjamanByUserID(ctx, userID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get list peminjaman by user ID: %v", err)
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	results := make([]payload.PeminjamanResponse, len(peminjamanList))
	for i, p := range peminjamanList {
		results[i] = payload.PeminjamanResponse{
			PeminjamanID:     p.PeminjamanID,
			NamaInventaris:   p.NamaInventaris,
			TanggalPinjam:    p.TglPinjam,
			TanggalKembali:   p.TglKembali,
			StatusPeminjaman: string(p.StatusPeminjaman), // .String() method
			Notes:            p.Notes.String,
		}
	}
	return results, nil
}

func (s *PeminjamanService) GetListPendingPeminjaman(ctx context.Context) ([]payload.PendingPeminjamanResponse, error) {
	q := sqlc.New(s.mainDB)
	pendingList, err := q.ListPendingPeminjaman(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get pending peminjaman list: %v", err)
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	results := make([]payload.PendingPeminjamanResponse, len(pendingList))
	for i, p := range pendingList {
		results[i] = payload.PendingPeminjamanResponse{
			PeminjamanID:   p.PeminjamanID,
			NamaInventaris: p.NamaInventaris,
			NamaPeminjam:   p.NamaPeminjamUser,
			UserIDPeminjam: p.UserIDPeminjam,
			TanggalPinjam:  p.TglPinjam,
			TanggalKembali: p.TglKembali,
			Notes:          p.Notes,
		}
	}
	return results, nil
}

func (s *PeminjamanService) GetListAvailableInventaris(ctx context.Context) ([]payload.AvailableInventarisResponse, error) {
	q := sqlc.New(s.mainDB)
	availableList, err := q.ListAvailableInventaris(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get available inventaris list: %v", err)
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	results := make([]payload.AvailableInventarisResponse, len(availableList))
	for i, inv := range availableList {

		status := ""
		if val, ok := inv.Status.(string); ok {
			status = val
		} else if val, ok := inv.Status.(sqlc.StatusPeminjamanEnum); ok {
			status = string(val)
		}

		results[i] = payload.AvailableInventarisResponse{
			InventarisID:   inv.InventarisID,
			NamaInventaris: inv.NamaInventaris,
			Status:         status,
			Keterangan:     inv.Keterangan,
		}
	}
	return results, nil
}

func (s *PeminjamanService) GetListOverduePeminjamanByUserID(ctx context.Context, userID string) ([]payload.OverduePeminjamanResponse, error) {
	q := sqlc.New(s.mainDB)
	overdueList, err := q.ListOverduePeminjamanByUserID(ctx, userID)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get overdue peminjaman list by user ID: %v", err)
		return nil, errors.WithStack(httpservice.ErrUnknownSource)
	}

	results := make([]payload.OverduePeminjamanResponse, len(overdueList))
	for i, op := range overdueList {
		results[i] = payload.OverduePeminjamanResponse{
			PeminjamanID:     op.PeminjamanID,
			NamaInventaris:   op.NamaInventaris,
			TanggalKembali:   op.TglKembali,
			TanggalPinjam:    op.TglPinjam,
			StatusPeminjaman: string(op.StatusPeminjaman),
		}
	}
	return results, nil
}
