package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"go-blog/models"
	"go-blog/models/db"
	_ "go-blog/routers"
)

var (
	syncdb = flag.Bool("syncdb", false, "\t同步数据库结构")
	force  = flag.Bool("force", false, "\tsyncdb 的强制模式，会丢失所有数据")
	run    = flag.Bool("run", false, "\t\t运行系统")
)

//初始化 beego 参数
func init() {
	//初始化日志引擎参数
	beego.SetLogger("file", `{"filename":"log/main.log"}`)
	if strings.ToLower(beego.RunMode) == "dev" {
		beego.SetLevel(beego.LevelDebug)
	} else {
		beego.SetLevel(beego.LevelNotice)
	}

	beego.SessionOn = true
	beego.SessionName = "go-blog-sessID"
}

func main() {
	if flag.Parse(); flag.NFlag() == 0 { //没有使用参数就打印帮助
		fmt.Println("命令行参数及默认值：")
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch {
	case *syncdb: //同步数据库结构
		verbose := strings.ToLower(beego.RunMode) == "dev"
		if err := orm.RunSyncdb("default", *force, verbose); err != nil {
			beego.Critical(err)
			os.Exit(1)
		}
		newAdminUser()
	case *run: //运行
		beego.Run()
	}

	time.Sleep(time.Second)
}

// 添加默认的管理员账号
func newAdminUser() {
	ac := &db.Accounts{
		Id:       1,
		Nickname: "Admin",
		Email:    "root@root.local",
		Password: models.Md5("12345678"),
		Admin:    true,
	}

	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(ac, "Id"); err != nil {
		beego.Critical(err)
		os.Exit(1)
	} else {
		if created {
			beego.Notice(`新安装，缺省用户："root@root.local" / "12345678"`)
		}
	}
}
