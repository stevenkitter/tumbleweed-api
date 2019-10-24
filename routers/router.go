package routers

import (

	"github.com/astaxie/beego"
	"github.com/stevenkitter/tumbleweed-api/controllers"
	"github.com/stevenkitter/tumbleweed-api/controllers/account"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    api := beego.NewNamespace("/api",
    	beego.NSNamespace("/user", beego.NSInclude(&account.UserController{})))
    beego.AddNamespace(api)
}
