package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.Data["mytest"] = "show it here"
	this.Data["title"] = "my packaging corp's"
	this.TplName = "index.html"
}
