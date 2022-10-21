package database

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ConnはDBとの接続情報をもつ構造体です
type Conn struct {
	DB *sqlx.DB
}

// NewConnはPostgreSQLを接続し、sql.DBオブジェクトのポインタをもつ構造体を返します
func NewConn() (*Conn, error) {
	dbDSN, err := config.DSN()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("postgres", dbDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL : %w", err)
	}

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to db ping : %w", err)
	}

	return &Conn{DB: db}, nil
}
