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
const defaultName = "default"

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
	ConnectName   string
	AutoMigrate   bool
}

var dbs = make(map[string]*DB) //保存连接列表
var eor error

func init() {
	dbs[defaultName] = new(DB)
}

func sql(user, password, addr, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, password, addr, dbname)
}

func Init(options Options) {
	db := new(DB)
	connectName := defaultName
	if !(len(dbs) == 0 || options.ConnectName == defaultName) {
		connectName = options.ConnectName
	}
	dbs[connectName] = db

	db.myDefault, eor = gorm.Open("mysql", sql(options.User, options.Password, options.Addr, options.Dbname))
	if eor != nil {
		panic(eor)
	} else {
		logs.Info(fmt.Sprintf("数据库[%s]连接成功", connectName))
	}

	db.myDefault.SingularTable(options.SingularTable)
	db.myDefault.LogMode(options.LogMode)
	if options.MaxIdle > 0 {
		db.myDefault.DB().SetMaxIdleConns(options.MaxIdle)
	}
	if options.MaxOpen > 0 {
		db.myDefault.DB().SetMaxOpenConns(options.MaxOpen)
	}

	e := db.myDefault.AutoMigrate(db.models...).Error
	if e != nil {
		logs.Error(e)
	}
}

func (db DB) Register(values ...interface{}) {
	db.models = append(db.models, values...)
}
func (db DB) MysqlNew() *gorm.DB {
	return db.myDefault.New()
}
func Default() *DB {
	return dbs[defaultName]
}

func Other(connectName string) (*DB) {
	if db, ok := dbs[connectName]; ok {
		return db
	}
	logs.Error(fmt.Sprintf("数据库连接[%s]不存在", connectName))
	return nil
}
