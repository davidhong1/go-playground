package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	sqlxDB *sqlx.DB
)

type SqlConfig struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	User            string `json:"user"`
	Pw              string `json:"pw"`
	Db              string `json:"db"`
	Charset         string `json:"charset"`
	MaxOpenConns    int    `json:"max_open_conns"`
	SetMaxIdleConns int    `json:"set_max_idle_conns"`
}

func (s SqlConfig) ConnectUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", s.User, s.Pw, s.Host, s.Port, s.Db, s.Charset)
}

func InitDB() error {
	sqlConfig := SqlConfig{
		Host:    "127.0.0.1",
		Port:    3306,
		User:    "root",
		Pw:      "sees7&chanting",
		Db:      "gostudydb",
		Charset: "utf8mb4",
	}

	var err error

	sqlxDB, err = sqlx.Connect("mysql", sqlConfig.ConnectUrl())
	if err != nil {
		return err
	}
	sqlxDB.SetMaxOpenConns(20)
	sqlxDB.SetMaxIdleConns(10)

	return nil
}

func CloseDB() error {
	if sqlxDB != nil {
		return sqlxDB.Close()
	}
	return nil
}
