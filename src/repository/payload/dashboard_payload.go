package payload

import (
	"database/sql"
	"time"
)

// DashboardStatsPayload representasi data statistik dashboard
type DashboardStatsPayload struct {
	TotalInventaris  int64 `json:"total_inventaris"`
	SedangDipinjam   int64 `json:"sedang_dipinjam"`
	Tersedia         int64 `json:"tersedia"`
	RusakMaintenance int64 `json:"rusak_maintenance"`
}

// RecentActivityPayload representasi aktivitas terbaru
type RecentActivityPayload struct {
	PeminjamanID   string         `json:"peminjaman_id"`
	NamaInventaris sql.NullString `json:"nama_inventaris"` // TETAP: Karena dari LEFT JOIN, bisa NULL
	TanggalPinjam  time.Time      `json:"tanggal_pinjam"`
	TanggalKembali time.Time      `json:"tanggal_kembali"` // UBAH: Jadi time.Time (asumsi dari sqlc ini NOT NULL)
	StatusDisplay  string         `json:"status_display"`
	NamaPeminjam   sql.NullString `json:"nama_peminjam"` // TETAP: Karena dari LEFT JOIN, bisa NULL
}

// RecentPeminjamPayload representasi peminjam terbaru
type RecentPeminjamPayload struct {
	UserID                string         `json:"user_id"`
	NamaPeminjam          sql.NullString `json:"nama_peminjam"`           // UBAH: Dibuat sql.NullString
	TanggalTerakhirPinjam time.Time      `json:"tanggal_terakhir_pinjam"` // TETAP: Karena MAX, asumsi NOT NULL
}

// NotReturnedInventarisPayload representasi inventaris belum dikembalikan
type NotReturnedInventarisPayload struct {
	PeminjamanID          string         `json:"peminjaman_id"`
	NamaInventaris        string         `json:"nama_inventaris"` // Diasumsikan selalu ada
	TanggalPinjam         time.Time      `json:"tanggal_pinjam"`
	TanggalKembaliRencana time.Time      `json:"tanggal_kembali_rencana"`
	NamaPeminjam          sql.NullString `json:"nama_peminjam"` // UBAH: Dibuat sql.NullString
}

// NewVendorPayload representasi vendor baru
type NewVendorPayload struct {
	VendorID     string         `json:"vendor_id"`
	NamaVendor   string         `json:"nama_vendor"`
	KontakVendor sql.NullString `json:"kontak_vendor"`
	JenisKontak  sql.NullString `json:"jenis_kontak"` // TETAP: Ini sudah benar sql.NullString di payload
	CreatedAt    time.Time      `json:"created_at"`
}

// DashboardResponse gabungan semua data dashboard
type DashboardResponse struct {
	Stats            DashboardStatsPayload          `json:"stats"`
	RecentActivities []RecentActivityPayload        `json:"recent_activities"`
	RecentPeminjam   []RecentPeminjamPayload        `json:"recent_peminjam"`
	NotReturned      []NotReturnedInventarisPayload `json:"not_returned"`
	NewVendors       []NewVendorPayload             `json:"new_vendors"`
}
