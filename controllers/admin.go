package controllers

/*
admin's page
can do this :
1. crud  categroy
2. crud  product
3. rd  inquery,feedback and reply
*/

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.Data["title"] = "Administrator Page"
	this.Data["IsAdmin"] = true
	this.TplNames = "admin.html"
}
func (this *AdminController) Post() {
	this.Data["title"] = "Administrator Page"
	this.Data["IsAdmin"] = true

	this.TplNames = "admin.html"

	/*
		//below code need add login page ,not here
		input := this.Input()                       // all post values
		input1 := this.Input().Get("key")           //get some one input by key
		this.Ctx.SetCookie("name", "value", 1, "/") //set the cookie,last 2 parameter is the maxage and path
		this.Redirect("/admin", 301)                //redirect to the admin page

		//-----login end
	*/

}

/*
func checkAccount(ctx *beego.Context) {

}
*/
