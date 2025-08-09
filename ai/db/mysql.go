package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"ai/config"
)

var DB *sql.DB

func InitMySQL() {
	var err error
	dsn := config.Getenv("DSN")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("MySQL连接失败: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("数据库不可用: %v", err)
	}
}