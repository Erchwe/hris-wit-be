package db

import (
	"time"

	"github.com/pkg/errors"
)

const (
	defaultMaxOpen           = 100
	defaultMaxLifetime       = 10 * time.Minute
	defaultMaxIdle           = 5
	defaultConnectTimeout    = 10 * time.Second
	defaultKeepAliveInterval = 30 * time.Second
)

// Option - database option.
type Option struct {
	Host          string
	Port          int
	Username      string
	Password      string
	DatabaseName  string
	ConnectionURL string
	*ConnectionOption
}

// ConnectionOption is db connection option.
type ConnectionOption struct {
	MaxIdle                int
	MaxLifetime            time.Duration
	MaxOpen                int
	ConnectTimeout         time.Duration
	KeepAliveCheckInterval time.Duration
}

// DefaultConnectionOption returns sensible conn setting.
func DefaultConnectionOption() *ConnectionOption {
	return &ConnectionOption{
		MaxIdle:                defaultMaxIdle,
		MaxOpen:                defaultMaxOpen,
		MaxLifetime:            defaultMaxLifetime,
		ConnectTimeout:         defaultConnectTimeout,
		KeepAliveCheckInterval: defaultKeepAliveInterval,
	}
}

var errInvalidDBSource = errors.New("invalid datasource host | port")

// NewDatabaseOption - default factory method.
func NewDatabaseOption(host string, port int, username, password, dbName string, conn *ConnectionOption) (*Option, error) {
	if host == "" || port == 0 {
		return nil, errors.Wrapf(errInvalidDBSource, "db: host=%s port=%d", host, port)
	}

	if conn == nil || conn.MaxOpen == 0 || conn.MaxOpen < conn.MaxIdle || conn.ConnectTimeout == 0 {
		conn = DefaultConnectionOption()
	}

	return &Option{
		Host:             host,
		Port:             port,
		Username:         username,
		Password:         password,
		DatabaseName:     dbName,
		ConnectionOption: conn,
	}, nil
}
