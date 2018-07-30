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

	//getEmailConfig()

	Log.Info("server start.")
	//select {}
}
