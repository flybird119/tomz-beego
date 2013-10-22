package controllers

import (
	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	// this.Data["user"] = NewDefaultUser()
	this.TplNames = "edit.html"
}
