package user

import (
	"github.com/astaxie/beego"
)


type RegController struct {
	beego.Controller
}

func (c *RegController) Get() {
	c.TplName = "user/reg.tpl"
}