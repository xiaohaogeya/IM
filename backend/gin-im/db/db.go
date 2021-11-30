package db

import (
	"fmt"
	"gin-im/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	mysqlUser := conf.AppConfig.Database.User
	mysqlPwd := conf.AppConfig.Database.Password
	mysqlPort := conf.AppConfig.Database.Port
	mysqldb := conf.AppConfig.Database.DbName
	mysqlHost := conf.AppConfig.Database.Host
	mysqlCharset := conf.AppConfig.Database.ChartSet
	mysqlSetMaxIdleConnNum := conf.AppConfig.Database.SetMaxIdleConnNum
	mysqlSetMaxOpenConnNum := conf.AppConfig.Database.SetMaxOpenConnNum

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		mysqlUser,
		mysqlPwd,
		mysqlHost,
		mysqlPort,
		mysqldb,
		mysqlCharset,
	)

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		DB = db
		sqlDB, _ := DB.DB()

		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(mysqlSetMaxIdleConnNum)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(mysqlSetMaxOpenConnNum)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
		log.Println("数据库连接成功")
	} else {
		panic("数据库连接失败")
	}
}
