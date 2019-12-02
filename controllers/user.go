package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}
func (this *UserController) Get() {
	this.Ctx.WriteString("Getinfo success")
}
func (this *UserController) GetInfo()  {
	id :=this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("GetInfo date,id="+id)
}
func (this *UserController) GetNum() {
	splate :=this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString("GetNum dat,username="+splate)
}
func (this *UserController)PostData()  {
	this.Ctx.WriteString("this is post func")

}