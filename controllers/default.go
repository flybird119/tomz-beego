package controllers

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
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

func GetAgeDesc(birth time.Time) string {
	temp := time.Now().Unix() - birth.Unix()
	t := time.Unix(int64(temp), 0)

	var month string
	if t.Month() >= 6 {
		month = "半"
	}

	return
}

func (this *MainController) Get() {
	this.Data["user"] = NewDefaultUser()
	this.TplNames = "index.html"
}
