package db

import (
	"Taichi/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB = nil

// 初始化数据库
func InitMySQL(mysql config.MySQL) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysql.Usr, mysql.Pwd, mysql.Addr, mysql.Port, mysql.Database)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	Db = db
	return nil
}
