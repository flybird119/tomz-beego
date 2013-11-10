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

	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.AddController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Run()
}
