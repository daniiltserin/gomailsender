package routers

import (
	"mailsender/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// beego.Router("/some", )
}
