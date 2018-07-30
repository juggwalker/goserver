package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func getEmailConfig() {
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM user").Values(&maps)
	for _, term := range maps {
		fmt.Println(term["id"], ":", term["name"])
	}
}
