package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sunlit-byte/blog/pkg/config"
	log "github.com/sunlit-byte/blog/pkg/logger"
)

var db *sql.DB

func InitDB() error {
	// 获取数据库连接配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	var err error

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("init mysql failed. err: %v", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("ping database failed. err: %v", err)
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}
