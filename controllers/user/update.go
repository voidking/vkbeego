package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"vkbeego/models"
)

type UpdateController struct {
	beego.Controller
}

func (c *UpdateController) Get() {
	c.Ctx.WriteString("Please use post request.")
}


func (c *UpdateController) Post() {
	//id,_ := c.GetInt("id")
	username := c.GetString("username")
	password := c.GetString("password")
	new_password := c.GetString("new_password")

	var o = orm.NewOrm()
	o.Using("default")

	exist := o.QueryTable("user").Filter("UserName", username).Exist()
	if !exist{
		c.Ctx.WriteString("Username doesn't exist!")
		return;
	}

	var user models.User
	o.QueryTable("user").Filter("Username",username).Filter("Password",password).One(&user)

	if o.Read(&user) == nil {
		user.Password = new_password
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
		}
	}else {
		c.Ctx.WriteString("Username or password is wrong!")
		return;
	}

	c.Ctx.WriteString("Password has been updated!")
}
