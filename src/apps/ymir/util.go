package main

import(
	log "github.com/sirupsen/logrus"
)

func LogIfError(err error) bool {
	if err != nil {
		log.Error(err)
		return true
	} else {
		return false
	}
}