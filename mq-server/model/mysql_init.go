package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	config "mq-server/conf"
	"os"
	"time"
)

// Db 数据库连接实例
var Db = &gorm.DB{}

func InitDb() {
	os.Setenv("GORM_LOG", "info")
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
	dbConfig := config.Cg.DbConfig
	dsn = fmt.Sprintf(dsn, dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败，请检查配置信息")
	}
	logger.Default.LogMode(logger.Info)
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	Db = db
}

func Test1() string {
	// 执行个简单的查询
	var result string
	Db.Raw("select version()").Scan(&result)
	fmt.Println(result)
	//Db.AutoMigrate(&User{})
	return ""
}
