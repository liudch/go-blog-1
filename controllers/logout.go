package controllers

import "github.com/astaxie/beego"

type LogoutController struct {
	BaseController
}

func (this *LogoutController) Get() {
	this.DelSession("Email")
	this.Redirect(beego.UrlFor("IndexController.GET"), 302)
	return
}
