package routers

import (
	"beego-tpl/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.AutoRouter(&controllers.ResourceController{})
}
