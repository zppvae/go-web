package controllers

import (
	"github.com/astaxie/beego"
	"go-web/models"
)

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

	topics, err := models.GetAllTopics(this.Input().Get("cate"), this.Input().Get("lable"), true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}