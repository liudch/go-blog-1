package main

import (
	"flag"
	"fmt"
	"os"

	_ "go-blog/models/db"
	_ "go-blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	syncdb = flag.Bool("syncdb", false, "\t同步数据库结构")
	force  = flag.Bool("force", false, "\tsyncdb 的强制模式，会丢失所有数据")
	run    = flag.Bool("run", true, "\t\t运行系统")
)

func main() {
	flag.Parse()
	if flag.NFlag() == 0 { //没有使用参数就打印帮助
		fmt.Println("命令行参数：")
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch {
	case *syncdb: //同步数据库结构
		err := orm.RunSyncdb("default", *force, true)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case *run: //运行
		beego.Run()
	}

	os.Exit(0)
}
