package controllers

import "github.com/oal/beego-pongo2"

type AccountController struct {
	BaseController
}

func (this *AccountController) Get() {
	pongo2.Render(this.Ctx, "account.html", pongo2.Context{
		"APP":   this.Data,
		"Title": "用户设置",
	})
}
