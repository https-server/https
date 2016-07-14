package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "net/http"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Upload() {
	if_switch := beego.AppConfig.String("switch")
	if if_switch == "on" {
		cfg_uname := beego.AppConfig.String("username")
		cfg_pwd := beego.AppConfig.String("password")

		r := this.Ctx.Request
		uname, pwd, ok := r.BasicAuth()

		if cfg_uname == uname && cfg_pwd == pwd {
			fmt.Println("-----", uname, pwd, ok, "-----")
		} else {
			fmt.Println("uname or pwd error")
			this.Redirect("/", 302)
			return
		}
	}

	_, h, err := this.GetFile("aaaa")
	if err != nil {
		beego.Error(err)
	}

	filename := h.Filename
	fmt.Println(filename)

	err = this.SaveToFile("aaaa", "download/"+filename)

	if err != nil {
		fmt.Println("Write file error: ", err)
		this.Redirect("/", 302)
		return
	}

	fmt.Println("Write file OK !!!")
	this.Redirect("/", 302)
}
