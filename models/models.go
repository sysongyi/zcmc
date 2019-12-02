package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)
type User struct {
	Id   int
	Name string `orm:size(100)`
	User_order []*User_order `orm:"reverse(many)"`
}
type User_order struct {
	Id int
	Order_data string `orm:size(100)`
	User *User `orm:"rel(fk)"`
}
func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:Sy123124@tcp(127.0.0.1:3306)/users?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User),new(User_order))

	// create table
	orm.RunSyncdb("default", false, true)
}
