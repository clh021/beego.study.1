package routers

import (
	"github.com/astaxie/beego"
	"go.study/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
