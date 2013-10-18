package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Title"] = "TomZhou"
	this.Data["Website"] = "tomz.me"
	this.Data["Email"] = "zhouytao@yeah.net"
	this.TplNames = "index.html"
}