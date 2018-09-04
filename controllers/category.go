package controllers

import (
	"github.com/astaxie/beego"
	"go-web/models"
	"github.com/astaxie/beego/validation"
	"log"
)

/*
   @Time : 2018/9/4 11:07 
   @Author : ff
*/


type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	// 检查是否有操作
	op := this.Input().Get("op")
	switch op {

	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	}

	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Categories"], err = models.GetCategorys()
	if err != nil {
		beego.Error(err)
	}
}

func (this *CategoryController) Post() {
	name := this.Input().Get("name")

	c := models.Category{Title: name}

	valid := validation.Validation{}
	valid.Required(c.Title, "title")
	valid.MaxSize(c.Title, 5, "titleMax")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	err := models.AddCategory(name)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/category", 302)
	return
}
