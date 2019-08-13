package routers

import (
	"vkbeego/controllers/user"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/login", &user.LoginController{})
	beego.Router("/user/reg", &user.RegController{})
	beego.Router("/user/add", &user.AddController{})
	beego.Router("/user/list", &user.ListController{})
	beego.Router("/user/update", &user.UpdateController{})
	beego.Router("/user/del", &user.DelController{})

}
