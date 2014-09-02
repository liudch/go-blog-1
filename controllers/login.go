package controllers

import (
	"github.com/oal/beego-pongo2"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	pongo2.Render(this.Ctx, "login.html", pongo2.Context{
		"APP":   this.Data,
		"Title": "请登陆！",
	})
}
