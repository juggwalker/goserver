package main

import (
	"github.com/go-ini/ini"
	"os"
)

func LoadConfig() map[string]interface{} {

	cfg, err := ini.Load("env.ini")
	if err != nil {
		Log.Fatal("Fail to read file: ", err)
		os.Exit(1)
	}

	ret := make(map[string]interface{})
	ret["app_mode"] = cfg.Section("").Key("app_mode").String()
	ret["server_protocol"] = cfg.Section("server").Key("protocol").In("http", []string{"http", "https"})
	ret["server_http_port"] = cfg.Section("server").Key("http_port").MustInt(12345)
	ret["mysql_host"] = cfg.Section("mysql").Key("host").String()
	ret["mysql_port"] = cfg.Section("mysql").Key("port").MustInt(3306)
	ret["mysql_dbname"] = cfg.Section("mysql").Key("dbname").String()
	ret["mysql_user"] = cfg.Section("mysql").Key("user").String()
	ret["mysql_pass"] = cfg.Section("mysql").Key("pass").String()
	ret["mysql_charset"] = cfg.Section("mysql").Key("charset").String()

	return ret
}
