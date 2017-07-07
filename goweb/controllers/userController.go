package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.Data["Website"] = "longdreamer"
	u.Data["Email"] = "astaxie@gmail.com"
	u.TplName = "index.tpl"
}
