package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Option struct {
	MaxIdleConn   int           // 最大空闲连接
	MaxOpenConn   int           // 最大连接数
	MaxLifeTime   int           // 空闲保活时间
	SlowThreshold time.Duration // 慢 SQL 阈值
}

var (
	db *gorm.DB
)

func NewOption() *Option {
	return &Option{
		MaxIdleConn:   10,
		MaxOpenConn:   100,
		MaxLifeTime:   7200,
		SlowThreshold: 500 * time.Microsecond,
	}
}

func GetDB() *gorm.DB {
	return db
}

func Init(dbType, dsn string, option *Option) error {
	var (
		err  error
		conf = &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: option.SlowThreshold,
					LogLevel:      logger.Warn,
					Colorful:      true,
				},
			),
		}
	)

	if dbType == "mssql" {
		db, err = gorm.Open(sqlserver.Open(dsn), conf)
	} else if dbType == "mysql" {
		db, err = gorm.Open(mysql.Open(dsn), conf)
	}

	dbPool, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbPool.SetMaxIdleConns(option.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	dbPool.SetMaxOpenConns(option.MaxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	dbPool.SetConnMaxLifetime(time.Duration(option.MaxLifeTime) * time.Second)

	if err != nil {
		return err
	}

	return nil
}
