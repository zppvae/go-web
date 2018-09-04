package controllers

import "github.com/astaxie/beego"

/*
   @Time : 2018/9/4 14:22 
   @Author : ff
*/

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.TplName = ""
}