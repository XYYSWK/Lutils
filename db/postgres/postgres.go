package postgres

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func PostgresInit(driverName, dataSourceName string, maxOpenConns, maxIdleConns int) *sql.DB {
	// 使用给定的驱动名称和数据源名称打开一个新的数据库连接
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Minute * 10)
	return db
}
