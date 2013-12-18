package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"tomz/models"
)

type MainController struct {
	beego.Controller
}

func NewDefaultUser() (u models.User) {
	return models.User{Name: "Tom Zhou", Sex: "M", Email: "zhouytao@yeah.net", HomePage: "tomz.blog"}
}

func GetAgeDesc(birth time.Time) string {
	temp := time.Now().Unix() - birth.Unix()
	t := time.Unix(int64(temp), 0)

	ageDesc := strconv.FormatInt(int64(t.Year()), 10) + "岁"
	if t.Month() >= 6 {
		ageDesc = ageDesc + "半"
	}

	return ageDesc
}

func (this *MainController) Get() {
	this.Data["user"] = NewDefaultUser()
	this.TplNames = "index.html"

	this.Data["isLogin"] = checkAccount(this.Ctx, this.Input())
}
