package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["IsMale"] = true
	c.Data["LuckyNumbers"] = [...]int{0,1,3,4}
	c.TplName = "test.tpl"
}

