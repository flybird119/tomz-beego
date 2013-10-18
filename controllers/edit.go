package controllers

import (
	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	this.Data["Website"] = "tomz.me"
	this.Data["Email"] = "zhouytao@yeah.net"
	this.TplNames = "edit.html"
}
