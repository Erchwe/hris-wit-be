package postgres

import (
	"database/sql"
	"fmt"
	"os" // <--- Import ini untuk membaca environment variables

	"github.com/wit-id/blueprint-backend-go/toolkit/config"
	"github.com/wit-id/blueprint-backend-go/toolkit/db"
)

func NewFromConfig(cfg config.KVStore, path string) (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")

	opt := &db.Option{}

	connOpt := db.DefaultConnectionOption()

	if maxIdle := cfg.GetInt(fmt.Sprintf("%s.conn.max-idle", path)); maxIdle > 0 {
		connOpt.MaxIdle = maxIdle
	}

	if maxOpen := cfg.GetInt(fmt.Sprintf("%s.conn.max-open", path)); maxOpen > 0 {
		connOpt.MaxOpen = maxOpen
	}

	if maxLifetime := cfg.GetDuration(fmt.Sprintf("%s.conn.max-lifetime", path)); maxLifetime > 0 {
		connOpt.MaxLifetime = maxLifetime
	}

	if connTimeout := cfg.GetDuration(fmt.Sprintf("%s.conn.timeout", path)); connTimeout > 0 {
		connOpt.ConnectTimeout = connTimeout
	}

	if keepAlive := cfg.GetDuration(fmt.Sprintf("%s.conn.keep-alive-interval", path)); keepAlive > 0 {
		connOpt.KeepAliveCheckInterval = keepAlive
	}

	opt.ConnectionOption = connOpt

	if databaseURL != "" {
		opt.ConnectionURL = databaseURL
	} else {

		opt.Host = cfg.GetString(fmt.Sprintf("%s.host", path))
		opt.Port = cfg.GetInt(fmt.Sprintf("%s.port", path))
		opt.Username = cfg.GetString(fmt.Sprintf("%s.username", path))
		opt.Password = cfg.GetString(fmt.Sprintf("%s.password", path))
		opt.DatabaseName = cfg.GetString(fmt.Sprintf("%s.schema", path))
	}

	return NewPostgresDatabase(opt)
}
