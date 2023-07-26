package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Model
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root/test?charset-utf8", 30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
	orm.RegisterDriver("mysql", orm.DRMySQL)
}

func main() {
	o := orm.NewOrm()
	user := User{Name: "wnn"}
	id, err := o.Insert(&user)
	user2 := User{Name: "wnn2"}
	id, err = o.Insert(&user2)
	fmt.Printf("id:%d,err:%v\n", id, err)

	user.Name = "wnnupdate"
	num, err := o.Update(&user)
	fmt.Printf("Num:%d, err:%v\n", num, err)

	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("err: %v\n", err)

	num, err = o.Delete(&u)
	fmt.Printf("num %d, err:%v\n", num, err)
}
