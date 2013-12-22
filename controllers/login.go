package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/url"
	"tomz/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	beego.Info("Login Get method")
	this.TplNames = "index.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if checkAccount(this.Ctx, this.Input()) {
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
	var username, password string
	ck, _ := ctx.Request.Cookie("username")
	if ck != nil {
		username = ck.Value
	}
	if len(username) == 0 {
		username = input.Get("username")
	}

	ck, _ = ctx.Request.Cookie("password")
	if ck != nil {
		password = ck.Value
	}
	if len(password) == 0 {
		password = input.Get("password")
	}

	u, err := models.GetById(1)
	if err != nil {
		panic("验证帐户信息出错！ " + err.Error())
	}

	if u.Email == username && u.PassWord == password {
		return true
	} else {
		return false
	}

	// return beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == passwrod
}
