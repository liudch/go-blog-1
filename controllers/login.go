package controllers

import (
	"github.com/dchest/captcha"
	"github.com/oal/beego-pongo2"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	Captcha := struct{ Id string }{captcha.NewLen(6)} //验证码长度为6

	pongo2.Render(this.Ctx, "login.html", pongo2.Context{
		"APP":     this.Data,
		"Title":   "请登陆！",
		"Captcha": &Captcha,
	})
}

func (this *LoginController) Post() {
	result := make(map[string]string)

	if !captcha.VerifyString(this.GetString("captchaId"), this.GetString("captchaVal")) {
		result["err"] = "验证码错误"
	}

	this.Data["json"] = result
	this.ServeJson()
}
