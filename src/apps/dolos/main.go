package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	utils Utils
	configs Configs
)

func main() {
	utils = Utils{}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	done := make(chan bool)

	conf, _ := newConfig()
	configs = conf

	//AzDiskFileIO{}.RunUseCase(done)

	newHttpServer()

	select {
		case <- done:
			return
	}
}