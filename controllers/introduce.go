package controllers

/*
news page
just show a page
*/

import (
	"github.com/astaxie/beego"
)

type IntroduceController struct {
	beego.Controller
}

func (this *IntroduceController) Get() {
	this.Data["IsAbout"] = true
	this.Data["title"] = "About"
	this.TplName = "aboutus.html"
}
