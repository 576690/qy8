package core

import (
	"fmt"
	"time"

	"github.com/576690/qy8/backend/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		//开发环境显示所有sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		//其他环境仅打印错误sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	//global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最大可容纳、
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //最大连接复用时间，不能超过mysql的wait_timeout
	return db
}
