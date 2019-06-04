/*
mysql.user = 用户名
mysql.password = 密码
mysql.host = ip地址
mysql.port = 端口
mysql.dbname = 数据库名字
gorm.singularTable = 全局禁用复数
gorm.logMode = 开启日志
gorm.maxOpen = 最大打开的连接数 0表示不限制
gorm.maxIdle = 闲置的连接数量
*/
package mysql

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//gorm model

type DB struct {
	myDefault *gorm.DB
	models    []interface{}
}

type Options struct {
	User          string
	Password      string
	Addr          string
	Dbname        string
	SingularTable bool
	LogMode       bool
	MaxIdle       int
	MaxOpen       int
}

var eor error
var db DB

func Init(options Options) {
	sql := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", options.User, options.Password, options.Addr, options.Dbname)
	db.myDefault, eor = gorm.Open("mysql", sql)
	if eor != nil {
		panic(eor)
	} else {
		logs.Info("mysql connect success")
	}

	db.myDefault.SingularTable(options.SingularTable)
	db.myDefault.LogMode(options.LogMode)
	if options.MaxIdle > 0 {
		db.myDefault.DB().SetMaxIdleConns(options.MaxIdle)
	}
	if options.MaxOpen > 0 {
		db.myDefault.DB().SetMaxOpenConns(options.MaxOpen)
	}
	e := db.MysqlNew().AutoMigrate(db.models...).Error
	if e != nil {
		fmt.Println(e)
	}
}

func (DB) Register(values ...interface{}) {
	db.models = append(db.models, values...)

}

func (DB) MysqlNew() *gorm.DB {
	if db.myDefault == nil {
		logs.Info("连接错误")
	}
	return db.myDefault.New()
}
