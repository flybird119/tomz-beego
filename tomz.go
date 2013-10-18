package main

import (
	"github.com/astaxie/beego"
	"tomz/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Run()
}
