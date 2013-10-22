package controllers

import (
	"github.com/astaxie/beego"
	"tomz/models"
)

type MainController struct {
	beego.Controller
}

func NewDefaultUser(count int) (u []models.User) {
	var users = [count]models.User
	for i := 0; i < count; i++ {
		users[i] = models.User{"Tom Zhou", 23, "M", "zhouytao@yeah.net", "tomz.blog"}
	}
	return users;
}

func (this *MainController) Get() {
	this.Data["user"] = NewDefaultUser()
	this.TplNames = "index.html"
}
