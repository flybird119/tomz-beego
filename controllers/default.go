package controllers

import (
	"github.com/astaxie/beego"
	"tomz/models"
)

type MainController struct {
	beego.Controller
}

func NewDefaultUser(count int) (u []models.User) {
	var users [int]models.User
	for i := 0; i < count; i++ {
		users[i] = models.User{Name: "Tom Zhou", Age: 23, Sex: "M", Email: "zhouytao@yeah.net", HomePage: "tomz.blog"}
	}
	return users
}

func (this *MainController) Get() {
	this.Data["user"] = NewDefaultUser()
	this.TplNames = "index.html"
}
