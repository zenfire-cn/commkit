package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type dbType struct {
	MYSQL string
	MSSQL string
}

type Option struct {
	Type        string // 数据库类型
	Host        string // 主机
	Port        string // 端口
	User        string // 用户名
	Password    string // 密码
	DbName      string // 连接的数据库名
	Args        string // 连接参数，如mysql常用的这些参数 "charset=utf8&parseTime=True&loc=Local"
	LogMode     bool   // 是否将sql语句输出到控制台
	MaxIdleConn int    // 最大空闲连接
	MaxOpenConn int    // 最大连接数
	MaxLifeTime int    // 连接超时时间
}

var (
	db     *gorm.DB
	DBType = &dbType{"mysql", "mssql"}
)

func DB() *gorm.DB {
	return db
}

func Init(option *Option) error {
	var (
		err     error
		connStr string
	)
	switch option.Type {
	case DBType.MSSQL:
		connStr = fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s",
			option.Host, option.Port, option.DbName, option.User, option.Password)
	default: // "mysql"
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			option.User, option.Password, option.Host, option.Port, option.DbName, option.Args)
	}
	db, err = gorm.Open(option.Type, connStr)
	if err != nil {
		return err
	}
	db.LogMode(option.LogMode)
	db.DB().SetMaxIdleConns(option.MaxIdleConn)
	db.DB().SetMaxOpenConns(option.MaxOpenConn)
	if option.MaxLifeTime == 0 {
		option.MaxLifeTime = 15 // 默认15秒超时
	}
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(option.MaxLifeTime)) // 建立连接的最大生命周期
	return nil
}

type txFunc func(db *gorm.DB) error

/**
 * @description: 事务操作的封装，像Tx函数中传入一个匿名函数，该匿名函数需要返回error对象，如果error为nil则提交，不为nil则回滚
 * @example:
	database.Tx(func(tx *gorm.DB) error {
		err := database.InsertUser()
		return err
	})
*/
func Tx(tf txFunc) error {
	tx := db.Begin()
	if err := tf(tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

/**
 * @description: 失败自动重连的实现
 * @params: -> option 初始化配置
 * @params: -> duration 重连间隔时间
 */
func InitUntilSuccess(option *Option, duration time.Duration) {
	// 失败重连
	if err := Init(option); err != nil {
		for {
			fmt.Println(err, "\nRe connect to the database......")
			time.Sleep(duration)
			if err = Init(option); err == nil {
				fmt.Println("Success re connect to the database.")
				break
			}
		}
	}
}
