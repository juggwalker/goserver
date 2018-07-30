package main

import (
	_ "fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()
var Config = LoadConfig()

func main() {
	Log.Out = os.Stdout

	defer func() {
		if r := recover(); r != nil {
			Log.Println("Recovered->", r)
		}
	}()

	go HttpServ()
	//go FileServ()

	//QuerySetting()

	//ll := (&FileInfo{FileName: "tttt.log"}).WriteWithIo("fsafjlkjfkajfkjkfkd\nfjksadjfklsajf\n22222222222222\n33333333333333333\n")
	//ll := (&FileInfo{FileName: "tttt.log"}).WriteWithBufio("fsafjlkjfkajfkjkfkd\nfjksadjfklsajf\n22222222222222\n33333333333333333\n")
	//fmt.Println(ll)

	Log.Info("server start.")
	select {}
}
