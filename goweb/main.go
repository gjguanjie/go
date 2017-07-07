package main

import (
	_ "goweb/configs"
	"goweb/models"
	_ "goweb/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.Info("配制文件：", beego.AppConfig.String("dev::httpport"))

	user := models.User{Name: "name", Email: "name@sina.com", Age: 10}
	num := models.SaveUser(&user)
	logs.Info("num:%d", num)
	beego.Run()
}
