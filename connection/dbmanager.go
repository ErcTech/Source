package erctech

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/BurntSushi/toml"
)

type DBManager struct {
	Database string
	User     string
	Password string
	Host     string
	Port     int
	PortREST int
}

var db *sql.DB

func PortREST() int {
	var ercDBManager DBManager
	if _, err := toml.DecodeFile("dbmanager.toml", &ercDBManager); err != nil {
		fmt.Println(err)
		return 0
	}

	return ercDBManager.PortREST
}

func GetDB() *sql.DB {
	if db == nil {
		var ercDBManager DBManager
		if _, err := toml.DecodeFile("dbmanager.toml", &ercDBManager); err != nil {
			fmt.Println(err)
			//			return db
		}

		strConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", ercDBManager.User, ercDBManager.Password, ercDBManager.Host, ercDBManager.Port, ercDBManager.Database)

		//fmt.Println(strConnectio)
		var err error
		db, err = sql.Open("mysql", strConnection)
		//err = db.Ping()

		if err != nil {
			panic(err.Error())
		}
	}

	return db
}
