package main

import (
	"os"
)

import (
	log "github.com/sirupsen/logrus"
)

var configs Configs = Configs{}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	conf := newConfig()
	configs = conf

	newHttpServer()

}