package account

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("Register", c.Register)
}

// @router /register [post]
func (c *UserController) Register() {

}