package routers

import (
	"go-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 注册 beego 路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})

	//topic路由
	beego.Router("/topic/view/:id", &controllers.TopicController{},"get:View")

	beego.AutoRouter(&controllers.TopicController{})

	//回复
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
}
