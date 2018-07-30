package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var dbconnection = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", Config["mysql_user"], Config["mysql_pass"], Config["mysql_host"], Config["mysql_port"], Config["mysql_dbname"], Config["mysql_charset"])

func getEmailConfig() {
	orm.RegisterDataBase("default", "mysql", dbconnection, 30)
	o := orm.NewOrm()
	var maps []orm.Params
	num, _ := o.Raw("SELECT id,_key,title FROM main_setting").Values(&maps)
	for _, term := range maps {
		fmt.Println(term["id"], ":", term["_key"], ":", term["title"])
	}
	fmt.Println(num)
}
