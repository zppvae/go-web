package controllers

import (
	"github.com/astaxie/beego"
	"go-web/models"
)

/*
   @Time : 2018/9/4 17:04 
   @Author : ff
*/

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid")
	err := models.AddReply(tid,
		this.Input().Get("nickname"), this.Input().Get("content"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	if !checkAccount(this.Ctx) {
		return
	}
	tid := this.Input().Get("tid")
	err := models.DeleteReply(this.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}