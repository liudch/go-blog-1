package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

type NextPreparer interface {
	NextPrepare()
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	// Setting properties.
	this.Data["Name"] = beego.AppName
	this.Data["Ver"] = beego.AppConfig.String("AppVer")
	this.Data["Url"] = strings.TrimRight(beego.AppConfig.String("AppUrl"), "/")
	this.Data["Admin"] = this.GetSession("Email") != nil //是否已经登陆

	if c, ok := this.AppController.(NextPreparer); ok {
		c.NextPrepare()
	}
}
