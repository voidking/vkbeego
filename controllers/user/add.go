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

	var o = orm.NewOrm()
	o.Using("default")
	exist := o.QueryTable("user").Filter("UserName", username).Exist()
	if exist{
		c.Data["json"] = map[string]interface{}{"code": 1, "ext": "Username has been existed!"}
		c.ServeJSON()
		return
	}

	var user = new(models.User)
	user.Username = username
	user.Password = password
	id,err := o.Insert(user)
	if err == nil {
		//fmt.Println(id)
		c.Data["json"] = map[string]interface{}{"code": 0, "userid": id,"ext": "User has been added!"}
		c.ServeJSON()
	}else {
		c.Data["json"] = map[string]interface{}{"code": 2, "ext": "Write to database failed!"}
		c.ServeJSON()
	}
	//c.Ctx.WriteString("add user: " + username)

	return
}
