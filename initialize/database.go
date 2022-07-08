package initialize

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局数据库 db
var db *gorm.DB

func init() {
	createDB()
}

func createDB() {
	fmt.Println("开始创建数据库连接。。。。。。。。。。。")
	dns := "root:root@tcp(127.0.0.1:3306)/goapp?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("创建数据库连接错误： " + err.Error())
		panic("数据库连接失败, error" + err.Error())
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("设置连接池错误： " + err.Error())
		panic("设置连接池错误, error" + err.Error())
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 设置连接空闲的最长时间。
	sqlDB.SetConnMaxIdleTime(time.Hour)
}

// 获取数据库连接
func MyDB() *gorm.DB {
	if db == nil {
		createDB()
	}
	return db
}
