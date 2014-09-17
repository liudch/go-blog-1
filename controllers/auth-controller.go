package controllers

type AuthController struct {
	BaseController
}

func (this *AuthController) NextPrepare() {
	//if this.GetSession("Email") != nil { //验证是否已经登陆
	//	this.Data["Admin"] = true
	//} else {
	//	this.Data["Admin"] = false
	//	this.Redirect(beego.UrlFor("IndexController.GET"), 302)
	//}

	this.SetSession("Email", "root@root.local")
	this.Data["Admin"] = true
}
