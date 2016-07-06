package main

import (
	"github.com/astaxie/beego"
	_ "http_server/routers"
)

func init() {
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSCertFile = "/home/fli/CA/new.ca"
	beego.BConfig.Listen.HTTPSKeyFile = "/home/fli/CA/new.key"
	beego.BConfig.Listen.HTTPSPort = 2443
}

func main() {
	beego.Run()
}
