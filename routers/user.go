package routers

import (
	"vkbeego/controllers/user"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/login", &user.LoginController{})
	beego.Router("/user/reg", &user.RegController{})
}
