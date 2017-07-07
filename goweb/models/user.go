package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id    int
	Name  string
	Email string
	Age   int
}

var (
	goUser   string = beego.AppConfig.String("mysqlUser")
	goPwd    string = beego.AppConfig.String("mysqlPwd")
	goDb     string = beego.AppConfig.String("mysqlDb")
	goMax, _        = strconv.Atoi(beego.AppConfig.String("mysqlMax"))
	goMin, _        = strconv.Atoi(beego.AppConfig.String("mysqlMin"))
	goConn   string = goUser + ":" + goPwd + "@/" + goDb + "?charset=utf8"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", goConn, goMax, goMin)
	orm.RegisterModelWithPrefix("tbl_", new(User))
	orm.RunSyncdb("default", true, true)

}

//保存单个User对象
func SaveUser(user *User) (id int64) {

	o := orm.NewOrm()
	o.Using("defalut")
	num, err := o.Insert(user)
	if err != nil {
		fmt.Printf("插入user失败，err:%s", err)
		return
	}
	return num
}

//保存多个User对象
func saveMultiUser(users []User) (num int64) {
	o := orm.NewOrm()
	num, err := o.InsertMulti(len(users), users)
	if err != nil {
		fmt.Printf("插入multiUser失败", err)
		return
	}
	return num

}
