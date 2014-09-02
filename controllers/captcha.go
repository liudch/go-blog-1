//获取验证码ID
package controllers

import "github.com/dchest/captcha"

type CaptchaController struct {
	BaseController
}

func (this *CaptchaController) Get() {
	Captcha := struct{ Id string }{captcha.NewLen(6)} //验证码长度为6

	this.Data["json"] = &Captcha
	this.ServeJson()
}
