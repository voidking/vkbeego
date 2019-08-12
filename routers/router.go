package routers

import (
	"vkbeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
}
