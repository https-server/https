package main

import (
	"github.com/astaxie/beego"
	_ "https/routers"
)

func init() {
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSCertFile = "conf/newreq.ca"
	beego.BConfig.Listen.HTTPSKeyFile = "conf/newreq.key"
	beego.BConfig.Listen.HTTPSPort = 2443
}

func main() {
	beego.Run()
}
