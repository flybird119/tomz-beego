package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"tomz/controllers"
	"tomz/models"
)

func init() {
	// models
}

func main() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.AddController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Run()
}
