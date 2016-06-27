package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "go.study"
	c.Data["Email"] = "clh021@gmail.com"
	c.TplName = "index.tpl"
}

func (this *UserController) Get() {
	this.Ctx.WriteString("this is a user page")
	// c.Data["Website"] = "go.study"
	// c.Data["Email"] = "clh021@gmail.com"
	// c.TplName = "user.tpl"
}
