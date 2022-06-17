package models

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

var DBCBS *sql.DB

func InitCBS() error {
	config := GetConfiguration()
	var err error
	connection_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.CBS.Host,
		config.CBS.Port,
		config.CBS.User,
		config.CBS.Password,
		config.CBS.Dbname)
	DBCBS, err = sql.Open("postgres", connection_string)
	
	if err = DBCBS.Ping(); err != nil {
		return err
	}

	DBCBS.SetMaxOpenConns(config.CBS.Max_open_conns)
	DBCBS.SetMaxIdleConns(config.CBS.Max_idle_conns)
	DBCBS.SetConnMaxLifetime(time.Duration(config.CBS.Conn_max_lifetime) * time.Second)
	return nil
}
