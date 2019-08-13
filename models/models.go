package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id      int
	Username    string
	Password     string
}

func create_table(){
	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := false

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "db.sqlite3")
	create_table()
}