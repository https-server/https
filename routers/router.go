package routers

import (
	"github.com/astaxie/beego"
	"https/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.UploadController{}, "post:Upload")
}
