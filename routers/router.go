package routers

import (
	"github.com/astaxie/beego"
	"myapp/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/category/Test", &controllers.TestController{})

	beego.Router("/introduce", &controllers.IntroduceController{})
	//beego.Router("/feedback", &controllers.IntroduceController{})
	beego.Router("/news", &controllers.NewsController{})
	beego.Router("/contactus", &controllers.ContactUsController{})
	beego.Router("/admin", &controllers.AdminController{})
}
