package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 定义数据库连接信息
const (
	username = "root"      // 数据库用户名
	password = "123456"    // 数据库密码
	host     = "127.0.0.1" // 数据库地址
	port     = 3306        // 数据库端口
	dbName   = "test_db"   // 要连接的数据库名
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close() // 关闭数据库连接
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	fmt.Println("数据库连接成功")

	db.SetMaxOpenConns(20)                  // 最大打开的连接数，默认0（无限制）
	db.SetMaxIdleConns(10)                  // 最大空闲连接数
	db.SetConnMaxLifetime(1 * time.Hour)    // 连接的最大存活时间
	db.SetConnMaxIdleTime(30 * time.Minute) // 连接的最大空闲时间
}
