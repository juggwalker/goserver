package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	Dsn         string
	Db          *sql.DB
	SettingInfo settingTB
}
type settingTB struct {
	id    int
	_key  sql.NullString
	title sql.NullString
}

/**
 * 只初始化一个sql.DB对象，并不会立即建立一个数据库的网络连接
 */
func LoadDB() {
	var err error
	dbw := DbWorker{
		//Dsn: "root:123456@tcp(localhost:3306)/db_name?charset=utf8mb4",
		Dsn: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", Config["mysql_user"], Config["mysql_pass"], Config["mysql_host"], Config["mysql_port"], Config["mysql_dbname"], Config["mysql_charset"]),
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer dbw.Db.Close()

	dbw.queryData()
}

func (dbw *DbWorker) queryDataPre() {
	dbw.SettingInfo = settingTB{}
}
func (dbw *DbWorker) queryData() {
	stmt, _ := dbw.Db.Prepare(`SELECT * From main_setting`)
	defer stmt.Close()

	dbw.queryDataPre()

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Printf("query data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&dbw.SettingInfo.id, &dbw.SettingInfo._key, &dbw.SettingInfo.title)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		/*
			if !dbw.SettingInfo._key.Valid {
				dbw.SettingInfo._key.String = ""
			}
			if !dbw.SettingInfo.title.Valid {
				dbw.SettingInfo.title.String = ""
			}*/
		fmt.Println("get data, id: ", dbw.SettingInfo.id, " key: ", dbw.SettingInfo._key.String, " title: ", dbw.SettingInfo.title.String)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
