package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Upload() {
	_, _, err := this.GetFile("aaa")
	if err != nil {
		beego.Error(err)
	}

	err = this.SaveToFile("aaa", "/home/fli/test")

	if err != nil {
		fmt.Println("Write file error: ", err)
		this.Redirect("/", 302)
		return
	}

	fmt.Println("Write file OK !!!")
}
