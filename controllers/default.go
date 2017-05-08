package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tolife.yuhanle.com"
	c.Data["Email"] = "doxxs@foxmail.com"
	c.TplName = "index.tpl"
}
