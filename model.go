package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var dbconnection = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", Config["mysql_user"], Config["mysql_pass"], Config["mysql_host"], Config["mysql_port"], Config["mysql_dbname"], Config["mysql_charset"])

func getEmailConfig() {
	var maps []orm.Params
	sql := "SELECT id,_key,title,_value,default_value FROM main_setting where _key IN('sender_format','send_mailer', 'mail_prefix', 'mail_host', 'mail_port', 'mail_account', 'mail_password', 'mail_timeout')"
	orm.RegisterDataBase("default", "mysql", dbconnection, 30)
	o := orm.NewOrm()

	num, _ := o.Raw(sql).Values(&maps)
	for _, term := range maps {
		fmt.Println(term["id"], ":", term["_key"], ":", term["title"])
	}
	fmt.Println(num)
	fmt.Println(maps)
}
