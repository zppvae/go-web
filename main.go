package main

import (
	_ "go-web/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Informational("App started.")
	beego.Run()

}

