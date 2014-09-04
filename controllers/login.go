package controllers

import (
	"go-blog/models/forms"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
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
	form := &forms.LoginForm{} //直接把表单解析到 struct
	if err := this.ParseForm(form); err != nil {
		beego.Error(err)
		return
	}

	//验证表单
	valid := validation.Validation{}
	b, err := valid.Valid(&forms.LoginFormValid{
		Email:      form.Email,
		Password:   form.Password,
		CaptchaId:  form.CaptchaId,
		CaptchaVal: form.CaptchaVal,
	})
	if err != nil {
		beego.Error(err)
		return
	}

	if !b {
		errMsg := valid.Errors[0].Field + ": " + valid.Errors[0].Message // 只回显第一个错误
		this.Data["json"] = map[string]string{"err": errMsg}
	} else {
		this.SetSession("Email", form.Email)
		this.Data["json"] = map[string]string{"ok": "true"}
	}

	this.ServeJson()
}
