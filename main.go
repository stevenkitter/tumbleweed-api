package main

import (
	"github.com/astaxie/beego"
	_ "github.com/stevenkitter/tumbleweed-api/routers"

)

func main() {
	beego.Run()
}

