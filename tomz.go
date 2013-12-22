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

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.SetLevel(beego.LevelInfo)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/Article", &controllers.ArticleController{})
	beego.Router("/Gallery", &controllers.GalleryController{})
	beego.Router("/About", &controllers.EditController{})
	beego.Router("/Login", &controllers.LoginController{})
	beego.Run()
}
