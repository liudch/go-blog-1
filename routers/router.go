package routers

import (
	"github.com/astaxie/beego"

	"go-blog/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
}
