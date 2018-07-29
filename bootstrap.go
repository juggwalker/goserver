package main

import (
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

	LoadDB()

	Log.Info("server start.")
	select {}
}
