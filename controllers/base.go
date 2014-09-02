package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
	//user models.User
}

func (this *BaseController) Prepare() {
	// Setting properties.
	this.Data["Name"] = beego.AppName
	this.Data["Ver"] = beego.AppConfig.String("AppVer")
	this.Data["Url"] = strings.TrimRight(beego.AppConfig.String("AppUrl"), "/")

	if c, ok := this.AppController.(NestPreparer); ok {
		c.NestPrepare()
	}
}
