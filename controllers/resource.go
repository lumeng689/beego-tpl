package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type ResourceController struct {
	beego.Controller
}

func (c ResourceController) URLMapping() {

	c.Mapping("/aa", c.Aa)
}

func (c *ResourceController) Aa() {

	// 获取body内容
	body := c.Ctx.Input.RequestBody

	fmt.Println("======" + string(body))

	v1 := c.Ctx.Input.Query("address")
	fmt.Println("======" + v1)

	// 获取参数，返回string类型, 如果参数不存在返回none作为默认值
	username := c.GetString("username", "none")

	m := make(map[string]string)
	m["username"] = username
	m["B"] = "2"

	// 固定的json数据
	c.Data["json"] = m

	c.ServeJSON()
}
