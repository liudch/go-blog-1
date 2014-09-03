package db

import (
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))

	// 注册数据库
	switch DBDriver := beego.AppConfig.String("DBDriver"); {
	case strings.Contains(DBDriver, "postgres"):
		orm.RegisterDriver("postgres", orm.DR_MySQL)
		orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("DBUri"))
	case strings.Contains(DBDriver, "mysql"):
		orm.RegisterDriver("mysql", orm.DR_MySQL)
		orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("DBUri"))
	default:
		orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
		orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("DBUri"))
	}

	// 设置 ORM 参数
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 100)
	orm.DefaultTimeLoc = time.UTC

	// beego 的运行在开发模式下，打开 orm 的 DEBUG 选项
	orm.Debug = beego.RunMode == "dev"
}
