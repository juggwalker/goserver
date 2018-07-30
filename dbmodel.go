package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var dbconnection1 = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", Config["mysql_user"], Config["mysql_pass"], Config["mysql_host"], Config["mysql_port"], Config["mysql_dbname"], Config["mysql_charset"])

func QuerySetting() {

	db, err := sql.Open("mysql", dbconnection1)
	if err != nil {
		panic(err)
		//return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,_key,title FROM main_setting")
	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {
		var id int
		var _key string
		var title string
		if err := rows.Scan(&id, &_key, &title); err != nil {
			Log.Warn(err)
		}
		fmt.Printf("%s id is %d onkey %s \n", title, id, _key)
	}

	if err := rows.Err(); err != nil {
		Log.Warn(err)
	}
	rows.Close()

}
