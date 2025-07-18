package echohttp

import (
	"context"
	application5 "github.com/wit-id/blueprint-backend-go/src/attachments/application"
	"github.com/wit-id/blueprint-backend-go/src/client/application"
	application3 "github.com/wit-id/blueprint-backend-go/src/lists/application"
	application2 "github.com/wit-id/blueprint-backend-go/src/projects/application"
	application4 "github.com/wit-id/blueprint-backend-go/src/tasks/application"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log" // <--- PASTIKAN INI DIIMPORT

	"github.com/wit-id/blueprint-backend-go/common/constants"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/echokit"

	"net/http"

	authTokenApp "github.com/wit-id/blueprint-backend-go/src/auth_token/application"
	authorizationBackofficeApp "github.com/wit-id/blueprint-backend-go/src/authorization/backoffice/application"
	authorizationHandheldApp "github.com/wit-id/blueprint-backend-go/src/authorization/handheld/application"
	emailApp "github.com/wit-id/blueprint-backend-go/src/email/application"

	userBackofficeApp "github.com/wit-id/blueprint-backend-go/src/user_backoffice/application"
	userBackofficeRoleApp "github.com/wit-id/blueprint-backend-go/src/user_backoffice_role/application"

	dataPegawaiApp "github.com/wit-id/blueprint-backend-go/src/data_pegawai/application"
	dataPelamarApp "github.com/wit-id/blueprint-backend-go/src/data_pelamar/application"
	lowonganPekerjaanApp "github.com/wit-id/blueprint-backend-go/src/lowongan_pekerjaan/application"
	userHandheldApp "github.com/wit-id/blueprint-backend-go/src/user_handheld/application"

	brandApp "github.com/wit-id/blueprint-backend-go/src/brand/application"
	damageHistoryApp "github.com/wit-id/blueprint-backend-go/src/damage_history/application"
	dashboardApp "github.com/wit-id/blueprint-backend-go/src/dashboard/application"
	inventarisApp "github.com/wit-id/blueprint-backend-go/src/inventaris/application"
	kantorApp "github.com/wit-id/blueprint-backend-go/src/kantor/application"
	kategoriApp "github.com/wit-id/blueprint-backend-go/src/kategori/application"
	kontakVendorApp "github.com/wit-id/blueprint-backend-go/src/kontak_vendor/application"
	peminjamanApp "github.com/wit-id/blueprint-backend-go/src/peminjaman/application"
	ruanganApp "github.com/wit-id/blueprint-backend-go/src/ruangan/application"
	usageHistoryApp "github.com/wit-id/blueprint-backend-go/src/usage_history/application"
	vendorKontakApp "github.com/wit-id/blueprint-backend-go/src/vendor_and_kontak/application"
	vendorApp "github.com/wit-id/blueprint-backend-go/src/vendors/application"
)

func RunEchoHTTPService(ctx context.Context, s *httpservice.Service, cfg config.KVStore) {
	log.Println("DEBUG: Starting RunEchoHTTPService function...") // <-- Log 1
	e := echo.New()
	e.HTTPErrorHandler = handleEchoError(cfg)

	e.Static("/uploads", "uploads")
	log.Println("DEBUG: Echo instance created, setting up CORS middleware.") // <-- Log 2

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://hris-wit-chernojovs-projects.vercel.app"}, // <--- UBAH INI
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, constants.DefaultAllowHeaderToken, constants.DefaultAllowHeaderRefreshToken},
	}))
	log.Println("DEBUG: CORS middleware configured.") // <-- Log 3

	runtimeCfg := echokit.NewRuntimeConfig(cfg, "restapi")
	runtimeCfg.HealthCheckFunc = s.GetServiceHealth

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	log.Println("DEBUG: Root / route hit.")        // <-- Log 4 (jika ada yang mengakses root)
	log.Println("DEBUG: Adding AuthToken routes.") // <-- Log 5

	authTokenApp.AddRouteAuthToken(s, cfg, e)
	log.Println("DEBUG: AuthToken routes added.") // <-- Log 6

	authorizationBackofficeApp.AddRouteAuthorizationBackoffice(s, cfg, e)
	authorizationHandheldApp.AddRouteAuthorizationHandheld(s, cfg, e)

	userBackofficeRoleApp.AddRouteUserBackofficeRole(s, cfg, e)
	userBackofficeApp.AddRouteUserBackoffice(s, cfg, e)

	userHandheldApp.AddRouteUserHandheld(s, cfg, e)

	application.AddRouteClient(s, cfg, e)
	application2.AddRouteProject(s, cfg, e)
	application3.AddRouteList(s, cfg, e)
	application4.AddRouteTask(s, cfg, e)
	application5.AddRouteAttachment(s, cfg, e)
	emailApp.AddRouteEmail(s, cfg, e)

	dataPelamarApp.AddRouteUserDataPelamar(s, cfg, e)
	lowonganPekerjaanApp.AddRouteLowonganPekerjaan(s, cfg, e)
	dataPegawaiApp.AddRouteDataPegawai(s, cfg, e)

	inventarisApp.AddRouteInventaris(s, cfg, e)
	brandApp.AddRouteBrand(s, cfg, e)
	damageHistoryApp.AddRouteDamageHistory(s, cfg, e)
	kantorApp.AddRouteKantor(s, cfg, e)
	vendorApp.AddRouteVendor(s, cfg, e)
	kategoriApp.AddRouteKategori(s, cfg, e)
	ruanganApp.AddRouteRuangan(s, cfg, e)
	kontakVendorApp.AddRouteKontakVendor(s, cfg, e)
	peminjamanApp.AddRoutePeminjaman(s, cfg, e)
	usageHistoryApp.AddRouteUsageHistory(s, cfg, e)
	vendorKontakApp.AddRouteVendorAndKontak(s, cfg, e)
	dashboardApp.AddRouteDashboard(s, cfg, e)
	log.Println("DEBUG: Setting route config.") // <-- Log 7

	// set config routes for role access
	httpservice.SetRouteConfig(ctx, s, cfg, e)
	log.Println("DEBUG: Route config set.")         // <-- Log 8
	log.Println("DEBUG: Attempting to run server.") // <-- Log 9

	// run actual server
	echokit.RunServerWithContext(ctx, e, runtimeCfg)
	log.Println("DEBUG: Server stopped.") // <-- Log 10 (Ini hanya akan terlihat jika server berhenti dengan graceful)

}
