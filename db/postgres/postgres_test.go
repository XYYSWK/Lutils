package postgres

import (
	"fmt"
	"testing"
)

func TestPostgresInit(t *testing.T) {
	//服务器要求使用 SSL 连接，则 sslmode=require
	//服务器不要求使用 SSL 连接，则 sslmode=disable
	db := PostgresInit("postgres", "user=postgres password=123456 host=192.168.239.128 port=5432 dbname=test sslmode=disable", 10, 5)
	if db != nil {
		fmt.Println("Success!")
	}
}
