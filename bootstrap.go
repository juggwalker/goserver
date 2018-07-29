package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()
var Config = LoadConfig()

func main() {
	Log.Out = os.Stdout

	go HttpServ()
	//go FileServ()

	Log.Info("server start.")
	select {}
}
