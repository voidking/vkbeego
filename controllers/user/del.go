package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DelController struct {
	beego.Controller
}

func (c *DelController) Get() {
	c.Ctx.WriteString("Please use post request.")
}

func (c *DelController) Post() {
	id,_ := c.GetInt("id")

	var o = orm.NewOrm()
	o.Using("default")

	num,err := o.QueryTable("user").Filter("Id",id).Delete()
	fmt.Printf("Returned Rows Num: %d, %s \n", num, err)
	if num != 0 {
		c.Ctx.WriteString("User has been deleted!")
	}else{
		c.Ctx.WriteString("User Id doesn't exist! Can't delete!")
	}

}
