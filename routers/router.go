package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"

	"go-blog/controllers"
)

func init() {
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/user-profile", &controllers.AccountController{})
	beego.Router("/bookmark", &controllers.BookmarkController{})
	beego.Router("/bookmark/del", &controllers.BookmarkController{}, "get:DelBookmark")
	beego.Router("/bookmark/switch", &controllers.BookmarkController{}, "get:Switch")

	// 验证码服务
	beego.Router("/captcha/request", &controllers.CaptchaController{})
	beego.Handler("/captcha/*.png", captcha.Server(240, 80)) //验证图片的宽和高(px)
}
