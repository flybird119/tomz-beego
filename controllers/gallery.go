package controllers

import (
	"github.com/astaxie/beego"
)

type GalleryController struct {
	beego.Controller
}

func (this *GalleryController) Get() {
	this.Data["isGalleryPage"] = true
	this.TplNames = "gallery.html"
}
