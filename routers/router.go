package routers

import (
	"zcmc/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserController{},"get:Get")
    beego.Router("/user/*",&controllers.UserController{},"get:GetNum;post:PostData")
}
