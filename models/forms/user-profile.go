package forms

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"go-blog/models"
	"go-blog/models/db"
)

//表单字段
type UserProfileForm struct {
	Email          string `form:"email"`
	Password       string `form:"password"`
	NewEmail       string `form:"newEmail"`
	NewPassword    string `form:"newPassword"`
	CfmNewPassword string `form:"cfmNewPassword"`
	Nickname       string `form:"nickname"`
	Motto          string `form:"motto"`
}

//表单验证
type UserProfileFormValid struct {
	Email          string `valid:"Required; Email; MaxSize(160)"` // Email字段需要符合邮箱格式，并且最大长度不能大于100个字符
	Password       string `valid:"Required; MinSize(8); MaxSize(64)"`
	NewEmail       string `valid:"MaxSize(160)"`
	NewPassword    string `valid:"MaxSize(64)"`
	CfmNewPassword string `valid:"MaxSize(64)"`
	Nickname       string `valid:"MaxSize(64)"`
	Motto          string `valid:"MaxSize(255)"`
}

//自定义验证
func (f *UserProfileFormValid) Valid(v *validation.Validation) {
	// 因 beego 的表单验证隐含不允许为空值
	if f.NewEmail != "" {
		v.Email(f.NewEmail, "NewEmail").Message("NewEmail: Must be a valid email address")
	}

	if f.NewPassword != "" && len(f.NewPassword) < 8 {
		v.SetError("NewPassword", "Minimum size is 8")
		return
	}
	// <--

	if f.NewPassword != f.CfmNewPassword {
		v.SetError("ConfirmNewPassword", "不正确")
		return
	}

	// 数据库验证步骤
	ac := &db.Accounts{Email: f.Email}
	o := orm.NewOrm()

	if err := o.Read(ac, "Email"); err != nil {
		v.SetError("Email", "找不到")
		return
	}

	if ac.Password != models.Md5(f.Password) {
		v.SetError("Password", "不正确")
		return
	}
}
