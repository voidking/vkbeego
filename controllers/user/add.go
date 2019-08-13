package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"vkbeego/models"
)

type AddController struct {
	beego.Controller
}

func (c *AddController) Get() {
	var username = c.GetString("username")
	fmt.Println(username)
	c.Ctx.WriteString("Please add user by post request.")
}


func (c *AddController) Post() {
	var username = c.GetString("username")
	var password = c.GetString("password")
	//fmt.Println(username)
	//fmt.Println(password)
	var o = orm.NewOrm()
	o.Using("default")
	var user = new(models.User)
	user.Username = username
	user.Password = password
	o.Insert(user)
	c.Ctx.WriteString("add user: " + username)

}
