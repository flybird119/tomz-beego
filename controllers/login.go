package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
	"net/url"
	"tomz/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "index.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == password {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("username", username, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")

	}

	this.Redirect("/", 301)
	return
}

// 验证Cookie中的帐户信息
func checkAccount(ctx *context.Context, input url.Values) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	passwrod := ck.Value

	var user models.User = models.User{
		Id:       1,
		UserCode: "username",
		PassWord: "password",
	}

	u, err := models.GetById(1)
	if u.UserCode == username && u.PassWord == passwrod {
		beego.Info("用户登陆成功!")
		log.Println("用户登陆成功!")
	}

	// return beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == passwrod
	return user.UserCode == username && user.PassWord == passwrod
}
