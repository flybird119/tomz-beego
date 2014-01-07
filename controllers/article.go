package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	this.Data["isArticlePage"] = true
	this.TplNames = "article.html"

	this.Data["isLogin"] = checkAccount(this.Ctx, this.Input())
}
