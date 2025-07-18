package service

import (
	"database/sql"
	vendorservice "github.com/wit-id/blueprint-backend-go/src/vendors/service"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
)

type VendorAndKontakService struct {
	mainDB    *sql.DB
	Cfg       config.KVStore
	vendorSvc *vendorservice.VendorService
}

func NewVendorAndKontakService(
	mainDB *sql.DB,
	cfg config.KVStore,
	vendorSvc *vendorservice.VendorService,
) *VendorAndKontakService {
	return &VendorAndKontakService{
		mainDB:    mainDB,
		Cfg:       cfg,
		vendorSvc: vendorSvc,
	}
}
