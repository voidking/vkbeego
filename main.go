package main

import (
	"github.com/astaxie/beego"
	_ "vkbeego/routers"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "vkbeego/models"
)


func main() {
	beego.Run()
}