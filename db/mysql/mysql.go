package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func MysqlInit(driverName, dataSourceName string, maxOpenConns, maxIdleConns int) *sql.DB {
	// 使用给定的驱动名称和数据源名称打开一个新的数据库连接
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		// 如果打开数据库连接时出现错误，立即触发 panic，并打印错误信息
		panic(err)
	}
	// 检查数据库连接是否可用，如果不可用则触发 panic，并打印错误信息
	if err = db.Ping(); err != nil {
		panic(err)
	}
	// 设置数据库连接池中的最大打开连接数
	db.SetMaxOpenConns(maxOpenConns)
	// 设置数据库连接池中的最大空闲连接数
	db.SetMaxIdleConns(maxIdleConns)
	// 设置连接的最大复用时间，超过该时间后连接将被关闭并重新创建
	db.SetConnMaxLifetime(time.Minute * 10)
	// 返回初始化后的数据库连接实例
	return db
}
