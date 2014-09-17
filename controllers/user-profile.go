package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/oal/beego-pongo2"

	"go-blog/models/db"
	"go-blog/models/forms"
)

type AccountController struct {
	AuthController
}

func (this *AccountController) Get() {
	o := orm.NewOrm()
	ac := &db.Accounts{Email: this.GetSession("Email").(string)}

	if err := o.Read(ac, "Email"); err != nil {
		beego.Error(err)
		return
	}

	pongo2.Render(this.Ctx, "user-profile.html", pongo2.Context{
		"APP":     this.Data,
		"Title":   "用户配置",
		"Account": ac,
	})
}

func (this *AccountController) Post() {
	form := &forms.UserProfileForm{} //直接把表单解析到 struct
	if err := this.ParseForm(form); err != nil {
		beego.Error(err)
		return
	}

	//验证表单
	valid := validation.Validation{}
	b, err := valid.Valid(&forms.UserProfileFormValid{
		Email:          form.Email,
		Password:       form.Password,
		NewEmail:       form.NewEmail,
		NewPassword:    form.NewPassword,
		CfmNewPassword: form.CfmNewPassword,
		Nickname:       form.Nickname,
		Motto:          form.Motto,
	})
	if err != nil {
		beego.Error(err)
		return
	}

	if !b {
		var errMsg string = valid.Errors[0].Field // 只回显第一个错误
		if errMsg != "" {
			errMsg += ": " + valid.Errors[0].Message
		} else {
			errMsg += valid.Errors[0].Message
		}

		this.Data["json"] = map[string]string{"err": errMsg}
	} else {
		o := orm.NewOrm()
		ac := &db.Accounts{Email: form.Email}

		if o.Read(ac, "Email") == nil {
			if form.NewEmail != "" {
				ac.Email = form.NewEmail
			}

			if form.NewPassword != "" {
				ac.Password = form.NewPassword
			}

			ac.Nickname, ac.Motto = form.Nickname, form.Motto

			if _, err := o.Update(ac); err != nil {
				beego.Error(err)
				return
			}
		}

		this.Data["json"] = map[string]string{"ok": "ok"}
	}

	this.ServeJson()
}
