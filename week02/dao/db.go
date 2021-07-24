package dao

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
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
	ParseTime       bool   `json:"parse_time"`
}

func (s SqlConfig) ConnectUrl() string {
	more := ""
	if s.ParseTime {
		more += "&parseTime=True"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s%s", s.User, s.Pw, s.Host, s.Port, s.Db, s.Charset, more)
}

func InitDB(configFile string) error {
	// get config from json
	bs, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	sqlConfig := SqlConfig{}
	err = json.Unmarshal(bs, &sqlConfig)
	if err != nil {
		return err
	}

	sqlxDB, err = sqlx.Connect("mysql", sqlConfig.ConnectUrl())
	if err != nil {
		return err
	}
	sqlxDB.SetMaxOpenConns(sqlConfig.MaxOpenConns)
	sqlxDB.SetMaxIdleConns(sqlConfig.SetMaxIdleConns)

	return nil
}

func CloseDB() error {
	if sqlxDB != nil {
		return sqlxDB.Close()
	}
	return nil
}
