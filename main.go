package main

import (
	"github.com/astaxie/beego"
	_ "https/routers"
)

func init() {
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSCertFile = "/home/fli/CA/newreq.ca"
	beego.BConfig.Listen.HTTPSKeyFile = "/home/fli/CA/newreq.key"
	beego.BConfig.Listen.HTTPSPort = 2443
}

func main() {
	beego.Run()
}
