package forms

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/dchest/captcha"

	"go-blog/models"
	"go-blog/models/db"
)

//表单字段
type LoginForm struct {
	Email      string `form:"email"`
	Password   string `form:"password"`
	CaptchaId  string `form:"captchaId"`
	CaptchaVal string `form:"captchaVal"`
}

//表单验证
type LoginFormValid struct {
	Email      string `valid:"Required; Email; MaxSize(160)"` // Email字段需要符合邮箱格式，并且最大长度不能大于100个字符
	Password   string `valid:"Required; MinSize(8); MaxSize(64)"`
	CaptchaId  string `valid:"Required"`
	CaptchaVal string `valid:"Required"`
}

//自定义验证
func (this *LoginFormValid) Valid(v *validation.Validation) {
	if strings.EqualFold(this.Email, this.Password) {
		v.SetError("Password", "不能与用户名相同")
		return
	}

	if !captcha.VerifyString(this.CaptchaId, this.CaptchaVal) {
		v.SetError("Captcha", "错误")
		return
	}

	// 数据库验证步骤
	ac := &db.Accounts{Email: this.Email}
	o := orm.NewOrm()

	if err := o.Read(ac, "Email"); err != nil {
		v.SetError("Email", "找不到")
		return
	}

	if ac.Password != models.Md5(this.Password) {
		v.SetError("Password", "不正确")
		return
	}
}
