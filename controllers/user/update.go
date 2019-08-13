package user

import (
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
		c.Data["json"] = map[string]interface{}{"code": 1, "ext": "Username doesn't exist!"}
		c.ServeJSON()
		return
	}

	var user models.User
	o.QueryTable("user").Filter("Username",username).Filter("Password",password).One(&user)

	if o.Read(&user) == nil {
		user.Password = new_password
		if num, err := o.Update(&user); err == nil {
			if num != 0 {
				c.Data["json"] = map[string]interface{}{"code": 0, "ext": "Password has been updated!"}
				c.ServeJSON()
			}
		}
	}else {
		c.Data["json"] = map[string]interface{}{"code": 2, "ext": "Username or password is wrong!"}
		c.ServeJSON()
	}

	return
}
