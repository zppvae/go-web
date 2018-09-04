package controllers

import "github.com/astaxie/beego"

/*
   @Time : 2018/9/4 10:44 
   @Author : ff
*/

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}