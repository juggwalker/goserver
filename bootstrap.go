package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func main() {
	log.Out = os.Stdout
	go FileServ()
	go HttpServ()
	log.Info("server start.")
	select {}
}
