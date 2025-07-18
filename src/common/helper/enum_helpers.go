package helper

import (
	"fmt"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
)

// ParseStatusPeminjamanEnumManual mengkonversi string ke tipe StatusPeminjamanEnum secara manual.
func ParseStatusPeminjamanEnumManual(s string) (sqlc.StatusPeminjamanEnum, error) {
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
