package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"tomz/controllers"
	"tomz/models"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/views", "views")

	beego.Router("/", &controllers.MainController{})
	beego.Router("/Article", &controllers.ArticleController{})
	beego.Router("/Gallery", &controllers.GalleryController{})
	beego.Router("/About", &controllers.EditController{})
	beego.Run()
}
