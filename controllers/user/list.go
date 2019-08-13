package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"vkbeego/models"
)

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {
	var o = orm.NewOrm()
	o.Using("default")
	var qs orm.QuerySeter
	qs = o.QueryTable("user")
	var users []*models.User
	//num, err := qs.All(&users)
	num, err := qs.All(&users,"Id", "Username")
	fmt.Printf("Returned Rows Num: %d, %s \n", num, err)
	fmt.Println(users[0].Username)
	c.Data["users"] = users
	c.TplName = "user/list.tpl"
}
