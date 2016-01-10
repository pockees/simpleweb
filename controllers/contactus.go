package controllers

/*
news page
just show a page
*/

import (
	"github.com/astaxie/beego"
)

type ContactUsController struct {
	beego.Controller
}

func (this *ContactUsController) Get() {
	this.Data["title"] = "Contact Us"
	this.Data["IsContact"] = true
	this.TplNames = "contactus.html"
}
