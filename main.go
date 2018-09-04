package main

import (
	_ "go-web/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go-web/models"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}


func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// create table
	orm.RunSyncdb("default", false, true)

	// 启动 beego
	beego.Run()
}

