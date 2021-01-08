package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	utils Utils
)

func main() {
	utils = Utils{}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	done := make(chan bool)

	AzDiskFileIO{}.RunUseCase(done)

	select {
		case <- done:
			return
	}
}