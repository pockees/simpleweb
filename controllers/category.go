package controllers

/*
product category page
this is important

show all the product by category
*/

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}
type TestController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["title"] = "Our Products"
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	//this.Ctx.WriteString("from category")
}

func (this *TestController) Get() {
	this.Ctx.WriteString("from Test category")
}
