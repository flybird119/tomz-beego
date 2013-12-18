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
	beego.BeeLogger.SetLogger("file", `{"filename":"logs/logs.log"}`)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/Article", &controllers.ArticleController{})
	beego.Router("/Gallery", &controllers.GalleryController{})
	beego.Router("/About", &controllers.EditController{})
	beego.Router("/Login", &controllers.LoginController{})
	beego.Run()
}
