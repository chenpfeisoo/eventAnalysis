package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"k8s.io/klog"
)

var (
	DB  *sql.DB
	err error
)

const (
	DriverName     = "mysql"
	DataSourceName = "root"
)

func InitDB() error {
	DB, err = sql.Open("mysql", "root:123456@tcp(localhost:8899)/test")
	if err != nil {
		klog.Error(err.Error())
	}
	err = DB.Ping()
	if err != nil {
		klog.Error(err.Error())
	}
	return err
}
