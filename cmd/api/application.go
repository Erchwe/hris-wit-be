package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	// --- Impor dari blueprint Anda ---
	"github.com/wit-id/blueprint-backend-go/common/echohttp"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/toolkit/db/postgres"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
	"github.com/wit-id/blueprint-backend-go/toolkit/runtimekit"

	// --- Impor library eksternal ---
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func main() {
	var err error
	var loggerIsSet bool // <-- PERUBAHAN 1: Tambahkan flag

	setDefaultTimezone()

	appContext, cancel := runtimekit.NewRuntimeContext()
	defer func() {
		cancel()
		if err != nil {
			// ================== PERUBAHAN 2 ==================
			// Periksa flag, bukan log.Get()
			if loggerIsSet {
				log.FromCtx(appContext).Error(err, "found error")
			} else {
				fmt.Printf("found error before logger was initialized: %v\n", err)
			}
			// ===============================================
		}
	}()

	// Membaca konfigurasi dari Environment Variables, bukan file.
	appConfig, err := envConfigVariable()
	if err != nil {
		return
	}

	mainDB, err := postgres.NewFromConfig(appConfig, "db")
	if err != nil {
		return
	}

	// Logger baru diinisialisasi di sini
	logger, err := log.NewFromConfig(appConfig, "log")
	if err != nil {
		return
	}
	logger.Set()
	loggerIsSet = true // <-- PERUBAHAN 3: Atur flag setelah logger siap

	svc := httpservice.NewService(mainDB, appConfig)

	// Menyesuaikan port untuk Render sebelum memanggil echohttp.
	port := os.Getenv("PORT")
	if port != "" {
		appConfig.Set("restapi.port", port)
		log.FromCtx(appContext).Info(fmt.Sprintf("Overriding restapi.port with environment PORT: %s", port))
	}

	// Jalankan HTTP Service menggunakan fungsi terpusat dari echohttp.
	echohttp.RunEchoHTTPService(appContext, svc, appConfig)
}

func setDefaultTimezone() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		loc = time.Now().Location()
	}
	time.Local = loc
}

func envConfigVariable() (cfg *viper.Viper, err error) {
	cfg = viper.New()
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()

	cfg.SetConfigFile("config.yaml")
	if readErr := cfg.ReadInConfig(); readErr != nil {
		// Menghapus panggilan log di sini karena logger belum diinisialisasi.
	}

	if os.Getenv("DATABASE_URL") == "" && os.Getenv("DB_HOST") == "" {
		err = errors.New("FATAL: database configuration not found in environment variables")
		return
	}

	return
}
