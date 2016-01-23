package controllers

/*
news page
just show message from
*/

import (
	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

func (this *NewsController) Get() {
	this.Data["IsNews"] = true
	this.Data["title"] = "News"
	this.TplName = "news.html"
}
