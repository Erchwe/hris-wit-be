package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
	"time"
)

type DashboardService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewDashboardService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *DashboardService {
	return &DashboardService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}

func (s *DashboardService) GetDashboardData(ctx context.Context) (payload.DashboardResponse, error) {
	q := sqlc.New(s.mainDB)
	var dashboardData payload.DashboardResponse

	stats, err := q.GetInventarisStats(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get inventaris stats: %v", err)
		return dashboardData, errors.WithStack(err)
	}
	dashboardData.Stats = payload.DashboardStatsPayload{
		TotalInventaris:  stats.TotalInventaris,
		SedangDipinjam:   stats.SedangDipinjam,
		Tersedia:         stats.Tersedia,
		RusakMaintenance: stats.RusakMaintenance,
	}

	recentActivities, err := q.ListRecentActivities(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get recent activities: %v", err)
		return dashboardData, errors.WithStack(err)
	}
	dashboardData.RecentActivities = make([]payload.RecentActivityPayload, len(recentActivities))
	for i, ra := range recentActivities {
		dashboardData.RecentActivities[i] = payload.RecentActivityPayload{
			PeminjamanID:   ra.PeminjamanID,
			NamaInventaris: ra.NamaInventaris,
			TanggalPinjam:  ra.TglPinjam,
			TanggalKembali: ra.TglKembali,
			StatusDisplay:  ra.StatusDisplay,
			NamaPeminjam:   ra.NamaPeminjamUser,
		}
	}

	recentPeminjam, err := q.ListRecentPeminjam(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get recent peminjam")
		return dashboardData, errors.WithStack(httpservice.ErrUnknownSource)
	}
	dashboardData.RecentPeminjam = make([]payload.RecentPeminjamPayload, len(recentPeminjam))
	for i, rp := range recentPeminjam {
		tanggalTerakhirPinjam := time.Time{}
		if val, ok := rp.TanggalTerakhirPinjam.(time.Time); ok {
			tanggalTerakhirPinjam = val
		}

		dashboardData.RecentPeminjam[i] = payload.RecentPeminjamPayload{
			UserID:                rp.UserID,
			NamaPeminjam:          rp.NamaPeminjam,
			TanggalTerakhirPinjam: tanggalTerakhirPinjam,
		}
	}

	notReturned, err := q.ListNotReturnedInventaris(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get not returned inventaris")
		return dashboardData, errors.WithStack(httpservice.ErrUnknownSource)
	}
	dashboardData.NotReturned = make([]payload.NotReturnedInventarisPayload, len(notReturned))
	for i, nr := range notReturned {
		dashboardData.NotReturned[i] = payload.NotReturnedInventarisPayload{
			PeminjamanID:          nr.PeminjamanID,
			NamaInventaris:        nr.NamaInventaris,
			TanggalPinjam:         nr.TglPinjam,
			TanggalKembaliRencana: nr.TanggalKembaliRencana,
			NamaPeminjam:          nr.NamaPeminjam,
		}
	}

	newVendors, err := q.ListNewVendors(ctx)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed to get new vendors")
		return dashboardData, errors.WithStack(httpservice.ErrUnknownSource)
	}
	dashboardData.NewVendors = make([]payload.NewVendorPayload, len(newVendors))
	for i, nv := range newVendors {
		jenisKontakPayload := sql.NullString{}
		if val, ok := nv.JenisKontak.(sql.NullString); ok {
			jenisKontakPayload = val
		} else if val, ok := nv.JenisKontak.(string); ok {
			jenisKontakPayload = sql.NullString{String: val, Valid: true}
		}

		dashboardData.NewVendors[i] = payload.NewVendorPayload{
			VendorID:     nv.VendorID,
			NamaVendor:   nv.NamaVendor,
			KontakVendor: nv.KontakVendor,
			JenisKontak:  jenisKontakPayload,
			CreatedAt:    nv.CreatedAt,
		}
	}

	return dashboardData, nil
}
