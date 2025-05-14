package connections

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
	"github.com/w0ikid/sqlite-test/internal/configs"
)

type DBconnector interface {
	Connect() (*sql.DB, error)
}

// ---- sqlite -----
type SQLiteConnector struct {
	DSN string
}

func (s SQLiteConnector) Connect() (*sql.DB, error) {
	return sql.Open("sqlite3", s.DSN)
}

// ---- postgres -----

type PostgresConnector struct {
	DSN string
}

func (p PostgresConnector) Connect() (*sql.DB, error) {
	return sql.Open("pgx", p.DSN)
}

// ---- fabric -----
func GetConnector(cfg configs.DBConfig) (DBconnector, error) {
	switch cfg.DriverName() {
	case "sqlite":
		return SQLiteConnector{DSN: cfg.DSN()}, nil
	case "postgres":
		return PostgresConnector{DSN: cfg.DSN()}, nil
	default:
		return nil, fmt.Errorf("unknown driver: %s", cfg.DriverName())
	}
}
