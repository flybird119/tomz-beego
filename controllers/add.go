package controllers

import "github.com/astaxie/beego"

type AddController struct {
	beego.Controller
}

func (this *AddController) Get() {
	this.TplNames = "add.html"
}
