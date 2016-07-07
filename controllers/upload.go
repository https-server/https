package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Upload() {
	_, h, err := this.GetFile("aaaa")
	if err != nil {
		beego.Error(err)
	}

	filename := h.Filename
	fmt.Println(filename)

	err = this.SaveToFile("aaaa", "/home/fli/bbbbb")

	if err != nil {
		fmt.Println("Write file error: ", err)
		this.Redirect("/", 302)
		return
	}

	fmt.Println("Write file OK !!!")
	this.Redirect("/", 302)
}
