package main

import (
	"github.com/astaxie/beego/orm"
	"zcmc/models"
	_ "zcmc/routers"
	"github.com/astaxie/beego"
)

func inserUser()  {
	o :=orm.NewOrm()
	user :=models.User{}
	user.Name = "sy3"
	id,err :=o.Insert(&user)
	if err !=nil {
		beego.Info("insert error")
		return
	}
	beego.Info("Insert success,id=",id)
}
func queryUser()  {
	o :=orm.NewOrm()
	user :=models.User{Id:3}
	err :=o.Read(&user)
	if err !=nil{
		beego.Info("Query is error")
		return
	}
	beego.Info("Query sucess,user=",user)
}
func updateUser()  {
	o :=orm.NewOrm()
	user :=models.User{
		Id:3,
		Name: "update3",
	}
	id,err :=o.Update(&user)
	if err !=nil {
		beego.Info("update error")
		return
	}
	beego.Info("Update success,id=",id)
}
func deleteUser()  {
	o :=orm.NewOrm()
	user :=models.User{
		Id:3,
	}
	id,err :=o.Delete(&user)
	if err !=nil {
		beego.Info("delete error")
		return
	}
	beego.Info("Delete success,id=",id)
}
func inserOrder()  {
	o :=orm.NewOrm()
	order :=models.User_order{}
	order.Order_data = "this is order-2"
	user :=models.User{Id:2}
	order.User=&user
	id,err :=o.Insert(&order)
	if err !=nil {
		beego.Info("insert error")
		return
	}
	beego.Info("Insert success,id=",id)
}

func main() {
	//inserUser()
	//queryUser()
	//updateUser()
	//deleteUser()
	//inserOrder()
	//queryOrder()
	beego.Run()
}

